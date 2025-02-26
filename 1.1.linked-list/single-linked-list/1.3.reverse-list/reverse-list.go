package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}


func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil

	current := head
	for current != nil {
		nextNode := current.Next
		current.Next = prev
		prev = current
		current = nextNode
	}
	return prev
}

func main() {
    node1 := &ListNode{Val: 1}
    node2 := &ListNode{Val: 2}
    node3 := &ListNode{Val: 3}
    node4 := &ListNode{Val: 4}
    node5 := &ListNode{Val: 5}

    node1.Next = node2
    node2.Next = node3
    node3.Next = node4
    node4.Next = node5

    newHead := reverseList(node1)
    
    for newHead != nil {
        fmt.Print(newHead.Val, " -> ")
        newHead = newHead.Next
    }
    fmt.Println("nil")
}