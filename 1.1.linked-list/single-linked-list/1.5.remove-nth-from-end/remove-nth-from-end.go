package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := &ListNode{Val: 0}
	temp.Next = head

	slow := temp
	fast := temp.Next

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return temp.Next
}

func main() {
    head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
    n := 2
    updatedHead := removeNthFromEnd(head, n)

    // Print updated linked list
    for updatedHead != nil {
        fmt.Print(updatedHead.Val, " -> ")
        updatedHead = updatedHead.Next
    }
    fmt.Println("nil")
}