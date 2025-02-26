# Merging Two Sorted Linked Lists in Go

## Overview
This program merges two sorted **singly linked lists** into a single sorted list. The function `mergeTwoLists` takes two linked lists as input and returns a new merged linked list in sorted order.

## Approach
The program follows these steps:
1. **Initialize a Dummy Node**: A dummy node helps simplify list operations by acting as a placeholder.
2. **Compare Elements**: Traverse both lists, linking the smaller node to the merged list.
3. **Attach Remaining Nodes**: If one list is exhausted, append the remaining nodes from the other list.
4. **Return the Merged List**: The merged sorted list is returned.

---

## Code Breakdown

### **1. ListNode Structure**
```go
 type ListNode struct {
     Val  int
     Next *ListNode
 }
```
This struct represents a single node in the linked list. Each node contains:
- `Val`: The integer value stored in the node.
- `Next`: A pointer to the next node in the list.

### **2. Merging Two Sorted Linked Lists**
```go
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    // Create a dummy node to simplify list operations
    dummy := &ListNode{}
    current := dummy

    // Compare nodes from both lists and merge
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            current.Next = list1
            list1 = list1.Next
        } else {
            current.Next = list2
            list2 = list2.Next
        }
        current = current.Next // Move to the next node
    }

    // Attach any remaining nodes from either list
    if list1 != nil {
        current.Next = list1
    } else {
        current.Next = list2
    }

    return dummy.Next // Return the head of the merged list
}
```
#### **Why Use a Dummy Node?**
- Simplifies handling of the head node since we don’t need special conditions.
- Eliminates the need to check if the merged list is empty at every step.
- The merged list starts at `dummy.Next`, avoiding extra conditional checks.

#### **Explanation**:
- **Dummy Node**: Serves as a starting point for the merged list.
- **Iterate Through Both Lists**: Compare values and append the smaller one.
- **Handle Remaining Nodes**: Attach leftover nodes once one list is exhausted.
- **Return `dummy.Next`**: The actual head of the merged list.

### **3. Printing the Linked List**
```go
func printList(head *ListNode) {
    for head != nil {
        fmt.Print(head.Val, " -> ")
        head = head.Next
    }
    fmt.Println("nil")
}
```
#### **Explanation**:
- Traverses the linked list, printing each node’s value.
- Stops when reaching the end (`nil`).

### **4. Main Function**
```go
func main() {
    // Example: list1 = [1, 3, 5], list2 = [2, 4, 6]
    list1 := &ListNode{1, &ListNode{3, &ListNode{5, nil}}}
    list2 := &ListNode{2, &ListNode{4, &ListNode{6, nil}}}

    mergedList := mergeTwoLists(list1, list2)
    printList(mergedList) // Output: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> nil
}
```
#### **Explanation**:
- Creates two sorted linked lists: `list1` and `list2`.
- Calls `mergeTwoLists` to merge them.
- Prints the merged list using `printList`.

---

## **Time Complexity Analysis**
- **Merging two linked lists** takes **O(n + m)**, where `n` and `m` are the lengths of `list1` and `list2`.
- **Appending nodes to the merged list** takes **O(1) per operation**.

### **Overall Complexity**: **O(n + m)** (optimal solution)

---

## **Advantages of This Approach**
✅ Uses a **dummy node** to simplify merging.  
✅ Time complexity is **optimal** at **O(n + m)**.  
✅ Works efficiently with **singly linked lists**.  
✅ **No extra memory allocation** beyond pointers.

---

## **Alternative Approaches**
1. **Recursion-Based Solution**
   - Uses recursion instead of iteration.
   - More readable but has extra function call overhead (**O(n + m) space**).

2. **Merging Using Additional Lists**
   - Create a new list and copy values manually.
   - Less efficient due to extra memory usage.

---

## **Conclusion**
This Go program efficiently merges two sorted linked lists into a single sorted list. Using a **dummy node and an iterative approach**, the solution remains **efficient and scalable**. 🚀

