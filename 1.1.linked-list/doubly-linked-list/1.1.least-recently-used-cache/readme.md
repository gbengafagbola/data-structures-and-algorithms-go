# LRU Cache Implementation in Go (Without `list` Package)

## Overview
This project implements a **Least Recently Used (LRU) Cache** in Go without using the `container/list` package. Instead, it manually implements a **doubly linked list** along with a **hash map**, ensuring `O(1)` time complexity for `get` and `put` operations.

## Data Structures Used
1. **Hash Map (`map[int]*Node`)**:
   - Stores key-node pairs for fast lookups in O(1) time.
   - Maps keys to nodes in the manually implemented doubly linked list.

2. **Doubly Linked List (Manually Implemented)**:
   - Maintains the usage order of elements.
   - Moves the most recently used elements to the front.
   - Evicts the least recently used (LRU) element when the cache reaches capacity.

## Implementation Details

### Node Struct
```go
// Represents each node in the doubly linked list
type Node struct {
    key   int
    value int
    prev  *Node
    next  *Node
}
```
- Each node contains a key-value pair.
- `prev` and `next` pointers allow efficient list traversal.

### LRUCache Struct
```go
// Defines the LRU Cache structure
type LRUCache struct {
    capacity int
    cache    map[int]*Node
    head     *Node
    tail     *Node
}
```
- `capacity`: Maximum number of elements the cache can hold.
- `cache`: A hash map storing pointers to nodes in the doubly linked list.
- `head`: A dummy node marking the start of the doubly linked list.
- `tail`: A dummy node marking the end of the doubly linked list.

### Constructor
```go
func NewLRUCache(capacity int) *LRUCache {
    lru := &LRUCache{
        capacity: capacity,
        cache:    make(map[int]*Node),
        head:     &Node{},
        tail:     &Node{},
    }
    // Connect head and tail to initialize the doubly linked list
    lru.head.next = lru.tail  // Head's next points to tail
    lru.tail.prev = lru.head  // Tail's prev points to head
    return lru
}
```
- Initializes the LRU Cache with a given capacity.
- Creates an empty hash map.
- Sets up `head` and `tail` dummy nodes and connects them so that `head.next` points to `tail` and `tail.prev` points to `head`. This forms an initially empty doubly linked list.

### Helper Functions

#### `remove(node *Node)`
```go
func (l *LRUCache) remove(node *Node) {
    node.prev.next = node.next  // Bypass node in the list
    node.next.prev = node.prev
}
```
- Removes a node from the doubly linked list by adjusting the `next` and `prev` pointers of its neighboring nodes.

#### `insertToFront(node *Node)`
```go
func (l *LRUCache) insertToFront(node *Node) {
    node.next = l.head.next       // Point node's next to current first element
    node.prev = l.head            // Point node's prev to head
    l.head.next.prev = node       // Update old first node's prev to new node
    l.head.next = node            // Update head's next to new node
}
```
- Moves a node to the front, marking it as the most recently used.
- The new node’s `next` is set to `head.next` (the current first node).
- The new node’s `prev` is set to `head`.
- The old first node’s `prev` is updated to point to the new node.
- `head.next` is updated to point to the new node.

### Get Method
```go
func (l *LRUCache) Get(key int) int {
    if node, found := l.cache[key]; found {
        l.remove(node)          // Remove from current position
        l.insertToFront(node)   // Move to the front
        return node.value
    }
    return -1
}
```
- If the key exists:
  - The node is removed from its current position.
  - The node is reinserted at the front to mark it as recently used.
  - The node's value is returned.
- If the key does not exist, return `-1`.

### Put Method
```go
func (l *LRUCache) Put(key int, value int) {
    if node, found := l.cache[key]; found {
        l.remove(node)          // Remove from current position
        node.value = value      // Update value
        l.insertToFront(node)   // Move to the front
        return
    }
    
    if len(l.cache) >= l.capacity {
        lruNode := l.tail.prev  // Get least recently used node (before tail)
        l.remove(lruNode)       // Remove from list
        delete(l.cache, lruNode.key)  // Remove from map
    }

    newNode := &Node{key: key, value: value}
    l.insertToFront(newNode)
    l.cache[key] = newNode
}
```
- If the key exists:
  - The node is removed from its current position.
  - Its value is updated.
  - It is reinserted at the front.
- If the key does not exist and the cache is full:
  - The least recently used node (just before `tail`) is removed.
  - The node is deleted from the map.
  - A new node is created and inserted at the front.
- The new node is added to the map.

### Example Usage
```go
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
```

## Time Complexity
- **Get: O(1)** → Hash map lookup and list update.
- **Put: O(1)** → Hash map insert/update and list manipulation.

## Summary
This implementation efficiently maintains an LRU cache without relying on Go's `list` package, manually implementing a **doubly linked list** for optimal performance. 🚀

