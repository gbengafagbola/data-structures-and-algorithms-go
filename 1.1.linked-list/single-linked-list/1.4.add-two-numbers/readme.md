# Adding Two Numbers in a Linked List

## Overview
This Go program defines a function `addTwoNumbers` that adds two numbers represented as linked lists, where each node contains a single digit. The digits are stored in reverse order, meaning the 1's place is at the head of the list.

## Implementation

```go
package main

import (
    "fmt"
)

type ListNode struct {
    Val  int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    ans := &ListNode{Val: 0}  // Dummy node to start the result list
    pointer := ans            // Pointer to traverse the result list
    carry := 0                // Stores carry-over value
    sum := 0                  // Stores sum of current digits

    list1 := l1               // Pointer to traverse l1
    list2 := l2               // Pointer to traverse l2

    for list1 != nil || list2 != nil {
        sum = carry           // Start with carry value
        if list1 != nil {
            sum = sum + list1.Val
            list1 = list1.Next
        }

        if list2 != nil {
            sum = sum + list2.Val 
            list2 = list2.Next
        }

        carry = sum / 10       // Compute new carry
        pointer.Next = &ListNode{Val: sum % 10} // Store current digit
        pointer = pointer.Next
    }
    if carry > 0 {
        pointer.Next = &ListNode{Val: carry} // Add final carry if needed
    }
    return ans.Next // Return the result list starting after the dummy node
}

func printList(node *ListNode) {
    for node != nil {
        fmt.Print(node.Val, " -> ")
        node = node.Next
    }
    fmt.Println("nil")
}

func main() {
    // Creating test case: 342 + 465
    l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}} // Represents 342
    l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}} // Represents 465

    result := addTwoNumbers(l1, l2)
    fmt.Print("Result: ")
    printList(result) // Expected output: 7 -> 0 -> 8 -> nil (807)
}
```

## Explanation
1. **ListNode Struct:**
   - Defines a singly linked list node with an integer value (`Val`) and a pointer to the next node (`Next`).

2. **Function Initialization:**
   - `ans`: A dummy node that helps in constructing the result list.
   - `pointer`: A reference to traverse the result list.
   - `carry`: Stores the carry-over when sum exceeds 9.
   - `sum`: Holds the temporary sum of digits.
   - `list1` and `list2`: Pointers to traverse `l1` and `l2`.

3. **Loop Execution:**
   - Runs while either `list1` or `list2` is not `nil`.
   - Adds `Val` from `list1` if it exists and moves to the next node.
   - Adds `Val` from `list2` if it exists and moves to the next node.
   - Computes the new `carry` and stores the remainder (`sum % 10`) as a new node.
   - Moves `pointer` forward.

4. **Final Carry Handling:**
   - If `carry > 0` after exiting the loop, a new node is added with `carry` value.

5. **Return Statement:**
   - The function returns `ans.Next` to exclude the dummy node from the final list.

## Visual Representation

### Example 1: Adding (342) + (465)

#### Input Lists (Stored in Reverse Order):
```
List 1: 2 -> 4 -> 3  (represents 342)
List 2: 5 -> 6 -> 4  (represents 465)
```

#### Step-by-step Computation:
| Step | List1 Val | List2 Val | Sum  | Carry | New Node |
|------|----------|----------|------|-------|----------|
| 1    | 2        | 5        | 7    | 0     | 7        |
| 2    | 4        | 6        | 10   | 1     | 0        |
| 3    | 3        | 4        | 8    | 0     | 8        |

#### Output List:
```
7 -> 0 -> 8  (represents 807)
```

## Complexity Analysis
- **Time Complexity:** `O(max(m, n))`, where `m` and `n` are the lengths of `l1` and `l2`.
- **Space Complexity:** `O(max(m, n))` for the new linked list.

## Conclusion
The `addTwoNumbers` function efficiently adds two numbers represented as reversed linked lists, considering digit-wise summation and carry handling. The provided test case in `main()` demonstrates how the function works in practice.

