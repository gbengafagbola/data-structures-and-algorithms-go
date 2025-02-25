# Merging Two Sorted Arrays Using a Linked List in Go

## Overview
This program merges two sorted integer arrays (`arr1` and `arr2`) into a single sorted **singly linked list**. It then converts this linked list back into an array and prints the result.

## Approach
The program follows these steps:
1. **Initialize a Linked List (`List`)**: The `List` structure maintains a reference to the `head` (first node) and `tail` (last node) of the linked list.
2. **Iterate through both arrays**: Compare elements from `arr1` and `arr2` and insert the smaller element into the linked list.
3. **Append remaining elements**: If any elements remain in `arr1` or `arr2`, add them to the linked list.
4. **Convert the linked list back to an array**: Traverse the linked list and store values in a slice.

---

## Code Breakdown

### **1. Node Structure**
```go
 type Node struct {
     value int
     next  *Node
 }
```
This struct represents a single node in the linked list. Each node contains:
- `value`: The integer value stored in the node.
- `next`: A pointer to the next node in the list.

### **2. List Structure**
```go
 type List struct {
     head *Node
     tail *Node
 }
```
The `List` struct maintains two pointers:
- `head`: Points to the first node in the list.
- `tail`: Points to the last node to allow efficient appends.

### **3. Merging Two Sorted Arrays**
```go
func (l *List) mergeSort(arr1, arr2 []int) []int {
    l1, l2 := 0, 0

    for l1 < len(arr1) && l2 < len(arr2) {
        Node1 := &Node{value: arr1[l1]}
        Node2 := &Node{value: arr2[l2]}

        if Node1.value < Node2.value {
            l.append(Node1)
            l1++
        } else {
            l.append(Node2)
            l2++
        }
    }

    // Append remaining elements
    for l1 < len(arr1) {
        l.append(&Node{value: arr1[l1]})
        l1++
    }

    for l2 < len(arr2) {
        l.append(&Node{value: arr2[l2]})
        l2++
    }

    return l.toSlice()
}
```
#### **Explanation**:
- **Two pointers `l1` and `l2`** traverse `arr1` and `arr2`, respectively.
- **Comparison of elements**: The smaller element is appended to the linked list.
- **Handling remaining elements**: If one array is exhausted, the remaining elements of the other array are appended.

### **4. Appending a Node to the Linked List**
```go
func (l *List) append(node *Node) {
    if l.head == nil {
        l.head = node
        l.tail = node
    } else {
        l.tail.next = node
        l.tail = node // Move the tail forward
    }
}
```
#### **Explanation**:
- If the list is empty (`head == nil`), the new node is set as both `head` and `tail`.
- Otherwise, the new node is linked to `tail.next`, and `tail` is updated to the new node.

### **5. Converting the Linked List to an Array**
```go
func (l *List) toSlice() []int {
    var result []int
    current := l.head
    for current != nil {
        result = append(result, current.value)
        current = current.next
    }
    return result
}
```
#### **Explanation**:
- Starts at `head` and traverses the list, collecting values into a slice.
- The resulting slice is returned.

### **6. Main Function**
```go
func main() {
    arr1 := []int{1, 3, 5, 7}
    arr2 := []int{1, 2, 4, 6}

    list := &List{} // Create an instance of List
    result := list.mergeSort(arr1, arr2)

    fmt.Println(result) // Output: [1, 1, 2, 3, 4, 5, 6, 7]
}
```
#### **Explanation**:
- Initializes `arr1` and `arr2` as sorted arrays.
- Creates an empty linked list.
- Calls `mergeSort` to merge the arrays into the list and convert it back to a slice.
- Prints the final sorted array.

---

## **Time Complexity Analysis**
- **Merging two sorted arrays** takes **O(n + m)**, where `n` and `m` are the lengths of `arr1` and `arr2`.
- **Appending nodes to the linked list** takes **O(1) per operation**.
- **Converting the linked list to an array** takes **O(n + m)**.

### **Overall Complexity**: **O(n + m)** (efficient solution)

---

## **Advantages of This Approach**
✅ Uses a **linked list** for efficient merging without extra space.  
✅ Time complexity is **optimal** at **O(n + m)**.  
✅ The `append()` function ensures **efficient insertions** at the tail.  
✅ **No additional sorting** is required since the input arrays are sorted.  

---

## **Alternative Approaches**
1. **Using a New Slice Instead of a Linked List**
   - Directly store merged elements in a slice.
   - Requires resizing and shifting in some cases.
   - Less efficient than a linked list for frequent insertions.

2. **Using a Min Heap (Priority Queue)**
   - Insert all elements into a min heap and extract them in sorted order.
   - More complex (`O(n log n)`) compared to this approach (`O(n + m)`).

---

## **Conclusion**
This Go program efficiently merges two sorted arrays into a sorted linked list and converts it back to an array. The use of a **linked list** ensures optimal insertions, making this approach **efficient and scalable**. 🚀

