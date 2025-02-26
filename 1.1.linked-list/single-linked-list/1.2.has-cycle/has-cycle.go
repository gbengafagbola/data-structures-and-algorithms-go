package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
func hasCycle(head *ListNode) bool {
	tortoise := head
	hare := head

	for hare != nil && hare.Next != nil {
		tortoise = tortoise.Next
		hare = hare.Next.Next

		if tortoise == hare {
			return true
		}
	}
	return false
}


func main() {
    node1 := &ListNode{Val: 1}
    node2 := &ListNode{Val: 2}
    node3 := &ListNode{Val: 3}
    node4 := &ListNode{Val: 4}

    node1.Next = node2
    node2.Next = node3
    node3.Next = node4
    node4.Next = node2 // Creating a cycle

    result := hasCycle(node1)
    fmt.Println(result) // Output: true
}