package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	// Example: list1 = [1, 3, 5], list2 = [2, 4, 6]
	list1 := &ListNode{1, &ListNode{3, &ListNode{5, nil}}}
	list2 := &ListNode{2, &ListNode{4, &ListNode{6, nil}}}

	mergedList := mergeTwoLists(list1, list2)
	printList(mergedList) // Output: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> nil
}
