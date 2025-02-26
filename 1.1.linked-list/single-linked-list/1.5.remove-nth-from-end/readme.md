# Remove Nth Node From End of Linked List

## Problem Statement
Given a singly linked list, remove the `n`th node from the end of the list and return its head.

---

## Code Implementation

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := &ListNode{Val: 0} // Step 1
	temp.Next = head

	slow := temp // Step 2
	fast := temp.Next

	for i := 0; i < n; i++ { // Step 3
		fast = fast.Next
	}

	for fast != nil { // Step 4
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next // Step 5

	return temp.Next // Step 6
}
```

---

## Step-by-Step Explanation

### **Step 1: Create a Temporary Dummy Node**
```go
temp := &ListNode{Val: 0}
temp.Next = head
```
We create a dummy node (`temp`) and set its `Next` pointer to the `head` of the linked list. This helps in handling edge cases where the head node itself needs to be removed.

**Visual Representation:**
```
Dummy → 1 → 2 → 3 → 4 → 5
```

### **Step 2: Initialize Two Pointers**
```go
slow := temp
fast := temp.Next
```
We define two pointers:
- `slow`: Starts at `temp`
- `fast`: Starts at `temp.Next`, which is the original `head`

**Visual Representation:**
```
slow (temp) → 0 → 1 → 2 → 3 → 4 → 5
                     ↑
                    fast
```

### **Step 3: Move `fast` Pointer `n` Steps Ahead**
```go
for i := 0; i < n; i++ {
	fast = fast.Next
}
```
This moves the `fast` pointer `n` steps ahead.

Example (`n = 2`):
```
slow (temp) → 0 → 1 → 2 → 3 → 4 → 5
                              ↑
                             fast
```

### **Step 4: Move Both Pointers Until `fast` Reaches the End**
```go
for fast != nil {
	fast = fast.Next
	slow = slow.Next
}
```
We move both pointers one step at a time until `fast` reaches `nil`.
At this point, `slow` will be at the node **before** the one that needs to be removed.

```
slow (to remove) → 0 → 1 → 2 → 3 → 4 → 5
                                       ↑
                                      fast (nil)
```

### **Step 5: Remove the Target Node**
```go
slow.Next = slow.Next.Next
```
We adjust the `Next` pointer of `slow` to skip the `n`th node from the end, effectively deleting it from the list.

```
slow (after remove) → 0 → 1 → 2 → 3 → 5
```

### **Step 6: Return the New Head**
```go
return temp.Next
```
Since `temp` was a dummy node, we return `temp.Next` as the new head of the list.

---

## Complexity Analysis
- **Time Complexity:** `O(L)`, where `L` is the length of the linked list (since we traverse the list twice at most).
- **Space Complexity:** `O(1)`, since we use only a few extra pointers.

---

## Example Walkthrough
### **Input:**
```
head = [1, 2, 3, 4, 5], n = 2
```
### **Output:**
```
[1, 2, 3, 5]
```

---

## Edge Cases Considered
1. **Removing the Head Node** (e.g., `[1,2]` with `n = 2`)
2. **Single Node List** (e.g., `[1]` with `n = 1`)
3. **Removing the Last Node** (e.g., `[1,2,3]` with `n = 1`)
4. **List of Length `n`** (Removing the only node in the list)

---

## Conclusion
This approach effectively removes the `n`th node from the end of a linked list using two pointers. The dummy node ensures that we can easily handle cases where the head node is removed.

