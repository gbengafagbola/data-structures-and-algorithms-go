# Reversing a Linked List

## Overview
This Go program defines a function `reverseList` that reverses a singly linked list. It iterates through the list, reversing the direction of the links between nodes.

## Implementation

```go
package main

type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode = nil
    current := head
    for current != nil {
        nextNode := current.Next
        current.Next = prev
        prev = current
        current = nextNode
    }
    return prev
}
```

## Explanation
1. **Initialization:**
   - `prev` is set to `nil` since the new tail (original head) should point to `nil`.
   - `current` starts at the head of the linked list.
2. **Reversal Process:**
   - The `Next` pointer of each node is redirected to point to `prev`.
   - `prev` moves forward to the current node.
   - `current` advances to the next node.
3. **Termination:**
   - When `current` reaches `nil`, `prev` holds the new head of the reversed list.

## Visual Representation

### Original List:
```
1 -> 2 -> 3 -> 4 -> 5 -> nil
```

### After Reversal:
```
nil <- 1 <- 2 <- 3 <- 4 <- 5
```

### Step-by-step Reversal:
| Step | Current Node | Next Node | Previous Node | List Status |
|------|-------------|-----------|---------------|-------------|
| 1    | 1           | 2         | nil           | `nil <- 1` |
| 2    | 2           | 3         | 1             | `nil <- 1 <- 2` |
| 3    | 3           | 4         | 2             | `nil <- 1 <- 2 <- 3` |
| 4    | 4           | 5         | 3             | `nil <- 1 <- 2 <- 3 <- 4` |
| 5    | 5           | nil       | 4             | `nil <- 1 <- 2 <- 3 <- 4 <- 5` |

## Example Usage

```go
func main() {
    node1 := &ListNode{Val: 1}
    node2 := &ListNode{Val: 2}
    node3 := &ListNode{Val: 3}
    node4 := &ListNode{Val: 4}
    node5 := &ListNode{Val: 5}

    node1.Next = node2
    node2.Next = node3
    node3.Next = node4
    node4.Next = node5

    newHead := reverseList(node1)
    
    for newHead != nil {
        fmt.Print(newHead.Val, " -> ")
        newHead = newHead.Next
    }
    fmt.Println("nil")
}
```

## Complexity Analysis
- **Time Complexity:** `O(n)`, since each node is visited once.
- **Space Complexity:** `O(1)`, as the reversal is done in place without extra space.

## Conclusion
The `reverseList` function efficiently reverses a singly linked list by modifying the pointers iteratively. This technique is useful in various linked list operations and is widely used in algorithmic problems.