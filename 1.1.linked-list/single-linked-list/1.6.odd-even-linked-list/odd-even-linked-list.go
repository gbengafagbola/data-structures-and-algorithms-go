package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	temp := head
	odd := temp
	even := odd.Next
	evenList := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenList
	return temp

}
