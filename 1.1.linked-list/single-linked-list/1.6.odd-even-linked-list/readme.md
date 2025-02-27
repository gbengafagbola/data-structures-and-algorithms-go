# Odd-Even Linked List Explanation

## Problem Statement
Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.

## Approach
We use two pointers:
1. `odd` - to track the last node in the odd indexed list.
2. `even` - to track the last node in the even indexed list.

We also maintain a reference to the head of the even list (`evenList`) so that we can connect the odd list to it at the end.

---

## Code Breakdown

```go
package main

// Definition for singly-linked list.
type ListNode struct {
    Val  int
    Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
```
- We define a `ListNode` struct with an integer value `Val` and a pointer `Next` to the next node.
- `oddEvenList` is a function that takes the head of the linked list and rearranges it as per the problem statement.

### Step 1: Handle Edge Cases

```go
    if head == nil {
        return head
    }
```
- If the linked list is empty (`head == nil`), return `head` as there is nothing to process.

### Step 2: Initialize Pointers

```go
    temp := head  // Store the head to return later
    odd := temp   // Pointer to track odd-positioned nodes
    even := odd.Next  // Pointer to track even-positioned nodes
    evenList := even  // Store the head of the even list to attach later
```
- `temp` stores the original `head` for returning the final result.
- `odd` starts at `head` (first node of the list).
- `even` starts at `odd.Next` (second node of the list).
- `evenList` keeps a reference to the even list's head for later connection.

### Step 3: Reordering Process

```go
    for even != nil && even.Next != nil {
```
- Iterate as long as `even` and `even.Next` are not `nil`, meaning there are still nodes to process.

```go
        odd.Next = even.Next  // Connect odd node to the next odd node
        odd = odd.Next  // Move odd pointer forward
```
- Set the next node of `odd` to skip `even` and point to `even.Next` (the next odd-indexed node).
- Move `odd` pointer forward.

```go
        even.Next = odd.Next  // Connect even node to the next even node
        even = even.Next  // Move even pointer forward
    }
```
- Set the next node of `even` to skip `odd` and point to `odd.Next` (the next even-indexed node).
- Move `even` pointer forward.

### Step 4: Connect Odd and Even Lists

```go
    odd.Next = evenList  // Attach even list to the end of the odd list
    return temp  // Return the new head
}
```
- After the loop, `odd` points to the last odd-indexed node. Connect it to `evenList` to attach the even-indexed nodes.
- Return `temp` (which still points to the modified head of the list).

---

## Visual Representation

### Example 1
#### Input:
```
1 -> 2 -> 3 -> 4 -> 5 -> NULL
```
#### Step-by-Step Process:
```
Initial:  1 -> 2 -> 3 -> 4 -> 5 -> NULL

Step 1: Move odd (1) -> (3), even (2) -> (4)
         1 -> 3 -> 4 -> 5 -> NULL
         2 -> 4 -> 5 -> NULL

Step 2: Move odd (3) -> (5), even (4) -> NULL
         1 -> 3 -> 5 -> NULL
         2 -> 4 -> NULL

Final:  1 -> 3 -> 5 -> 2 -> 4 -> NULL
```
#### Output:
```
1 -> 3 -> 5 -> 2 -> 4 -> NULL
```

### Example 2
#### Input:
```
2 -> 1 -> 3 -> 5 -> 6 -> 4 -> 7 -> NULL
```
#### Output:
```
2 -> 3 -> 6 -> 7 -> 1 -> 5 -> 4 -> NULL
```

---

## Time & Space Complexity Analysis
- **Time Complexity**: **O(N)** (since we traverse the list once, separating odd and even nodes)
- **Space Complexity**: **O(1)** (since we modify the list in place without extra storage)

---

## Edge Cases Considered
1. **Empty List** (`head == nil`) → Return `nil`.
2. **Single Node** (`head.Next == nil`) → No rearrangement needed, return as is.
3. **Two Nodes** (`head.Next.Next == nil`) → Already in correct order.
4. **Even/Odd Length Lists** → Handled properly by iteration.

---

## Summary
- The function correctly separates odd and even indexed nodes while preserving relative order.
- Uses two pointers (`odd` and `even`) with `O(1)` space complexity.
- Efficiently modifies the linked list in a single pass with `O(N)` time complexity.
- Properly handles various edge cases.

