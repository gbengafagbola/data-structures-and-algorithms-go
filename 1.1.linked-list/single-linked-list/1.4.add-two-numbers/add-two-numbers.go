package main

type ListNode struct {
	Val int
	Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
		ans := &ListNode{Val: 0}

		pointer := ans
		carry := 0
		sum := 0

		list1 := l1
		list2 := l2

		for list1 != nil || list2 != nil {
			sum = carry
			if list1 != nil {
				sum = sum + list1.Val
				list1 = list1.Next
			}

			if list2 != nil {
				sum = sum + list2.Val 
				list2 = list2.Next
			}

			carry = sum/10
			pointer.Next = &ListNode{Val: sum%10}
			pointer = pointer.Next
		}
		if carry > 0 {
			pointer.Next = &ListNode{Val: carry}
		}
		return ans.Next
}