# Detecting a Cycle in a Linked List

## Overview
This Go program defines a function `hasCycle` that detects whether a given singly linked list contains a cycle. It uses Floyd's Cycle Detection Algorithm (also known as the Tortoise and Hare Algorithm), which employs two pointers moving at different speeds to detect a loop.

## Implementation

```go
package main

type ListNode struct {
    Val  int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
    tortoise := head
    hare := head

    for hare != nil && hare.Next != nil {
        tortoise = tortoise.Next
        hare = hare.Next.Next

        if tortoise == hare {
            return true
        }
    }
    return false
}
```

## Explanation
1. **Initialization:** Two pointers, `tortoise` and `hare`, both start at the head of the linked list.
2. **Traversal:** The `tortoise` moves one step at a time, while the `hare` moves two steps at a time.
3. **Cycle Detection:** If there is a cycle, the `hare` will eventually meet the `tortoise` inside the loop, returning `true`.
4. **Termination:** If `hare` reaches the end (`nil`), the function returns `false`, indicating no cycle.

## Example Usage

### Example 1: Linked List with a Cycle
```
Input: 1 -> 2 -> 3 -> 4 -> 2 (cycle)
Output: true
```
**Explanation:** The last node connects back to the second node, forming a cycle.

### Example 2: Linked List without a Cycle
```
Input: 1 -> 2 -> 3 -> 4 -> nil
Output: false
```
**Explanation:** The list terminates properly, so no cycle exists.

## Complexity Analysis
- **Time Complexity:** `O(n)`, where `n` is the number of nodes in the linked list. The `hare` and `tortoise` pointers traverse at most `n` steps.
- **Space Complexity:** `O(1)`, since only two pointers are used regardless of the input size.

## Usage
To use this function in your Go project, define your linked list structure and call `hasCycle` with the head node as input.

```go
func main() {
    node1 := &ListNode{Val: 1}
    node2 := &ListNode{Val: 2}
    node3 := &ListNode{Val: 3}
    node4 := &ListNode{Val: 4}

    node1.Next = node2
    node2.Next = node3
    node3.Next = node4
    node4.Next = node2 // Creating a cycle

    result := hasCycle(node1)
    fmt.Println(result) // Output: true
}
```

## Conclusion
The `hasCycle` function efficiently detects cycles in a linked list using Floyd’s Tortoise and Hare algorithm. It is widely used in linked list problems to identify loops efficiently.

