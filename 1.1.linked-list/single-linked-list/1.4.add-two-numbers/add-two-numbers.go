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