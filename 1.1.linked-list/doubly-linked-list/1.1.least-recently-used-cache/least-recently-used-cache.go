package main

import "fmt"

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func NewLRUCache(capacity int) *LRUCache {
	lru := &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     &Node{},
		tail:     &Node{},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (l *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) insertToFront(node *Node) {
	node.next = l.head.next
	node.prev = l.head
	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache) Get(key int) int {
	if node, found := l.cache[key]; found {
		l.remove(node)
		l.insertToFront(node)
		return node.value
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if node, found := l.cache[key]; found {
		l.remove(node)
		node.value = value
		l.insertToFront(node)
		return
	}

	if len(l.cache) >= l.capacity {
		lruNode := l.tail.prev
		l.remove(lruNode)
		delete(l.cache, lruNode.key)
	}

	newNode := &Node{key: key, value: value}
	l.insertToFront(newNode)
	l.cache[key] = newNode
}

func main() {
	cache := NewLRUCache(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // Output: 1
	cache.Put(3, 3) // Evicts key 2
	fmt.Println(cache.Get(2)) // Output: -1 (not found)
	cache.Put(4, 4) // Evicts key 1
	fmt.Println(cache.Get(1)) // Output: -1 (not found)
	fmt.Println(cache.Get(3)) // Output: 3
	fmt.Println(cache.Get(4)) // Output: 4
}
