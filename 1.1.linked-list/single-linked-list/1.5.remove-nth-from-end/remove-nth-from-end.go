package main

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