package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	order    *list.List
}

type entry struct {
	key   int
	value int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		order:    list.New(),
	}
}

func (l *LRUCache) Get(key int) int {
	if elem, found := l.cache[key]; found {
		l.order.MoveToFront(elem)
		return elem.Value.(*entry).value
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if elem, found := l.cache[key]; found {
		l.order.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}
	
	if len(l.cache) >= l.capacity {
		oldest := l.order.Back()
		if oldest != nil {
			delete(l.cache, oldest.Value.(*entry).key)
			l.order.Remove(oldest)
		}
	}

	newEntry := &entry{key, value}
	elem := l.order.PushFront(newEntry)
	l.cache[key] = elem
}

func main() {
	cache := NewLRUCache(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // 1
	cache.Put(3, 3)           // Evicts key 2
	fmt.Println(cache.Get(2)) // -1 (not found)
	cache.Put(4, 4)           // Evicts key 1
	fmt.Println(cache.Get(1)) // -1 (not found)
	fmt.Println(cache.Get(3)) // 3
	fmt.Println(cache.Get(4)) // 4
}
