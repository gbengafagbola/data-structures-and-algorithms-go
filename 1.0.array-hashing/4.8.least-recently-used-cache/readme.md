# LRU Cache Implementation in Go

## Overview
This project implements a **Least Recently Used (LRU) Cache** in Go using a combination of a **hash map** and a **doubly linked list** from the `container/list` package. This approach ensures that both `get` and `put` operations run in **O(1) time complexity**.

## Data Structures Used
1. **Hash Map (`map[int]*list.Element`)**:
   - Stores key-value pairs for quick lookups in O(1) time.
   - Maps keys to elements in the doubly linked list.

2. **Doubly Linked List (`list.List`)**:
   - Maintains the usage order of elements.
   - Moves the most recently used elements to the front.
   - Evicts the least recently used (LRU) element when the cache reaches capacity.

## Implementation Details

### LRUCache Struct
```go
// Defines the LRU Cache structure
type LRUCache struct {
    capacity int
    cache    map[int]*list.Element
    order    *list.List
}
```
- `capacity`: The maximum number of elements the cache can hold.
- `cache`: A hash map storing pointers to nodes in the doubly linked list.
- `order`: A doubly linked list tracking the most recently used elements.

### Entry Struct
```go
type entry struct {
    key   int
    value int
}
```
- Represents each cache entry with a key-value pair.

### Constructor
```go
func NewLRUCache(capacity int) *LRUCache {
    return &LRUCache{
        capacity: capacity,
        cache:    make(map[int]*list.Element),
        order:    list.New(),
    }
}
```
- Initializes the LRU Cache with a given capacity.
- Creates an empty hash map and doubly linked list.

### Get Method
```go
func (l *LRUCache) Get(key int) int {
    if elem, found := l.cache[key]; found {
        l.order.MoveToFront(elem)
        return elem.Value.(*entry).value
    }
    return -1
}
```
- If the key exists, move it to the front (most recently used) and return the value.
- Otherwise, return `-1` (not found).

### Put Method
```go
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
```
- If the key exists, move it to the front and update its value.
- If the cache exceeds capacity, evict the least recently used (LRU) entry.
- Insert the new key-value pair at the front.

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
This implementation efficiently maintains an LRU cache using a combination of **hash map** and **doubly linked list**, ensuring fast lookups and updates. 🚀

