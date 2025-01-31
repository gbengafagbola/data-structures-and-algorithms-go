package main

import (
	"fmt"
)

type List struct {
	head   *Node
	tail   *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Push(value string) {
	node := &Node{value: value}

	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
}

func (l *List) Pop() {
	l.tail = l.tail.prev
	l.tail.next = nil
}

func (l *List) Find(value string) (*Node) {
	node := l.First()

	for node.next != nil {
		if (node.value == value) {
			return node
		}
		node = node.next
	}

	return &Node{value: "Not found"}
}

type Node struct {
	value string
	next  *Node
	prev  *Node
}

// a method to return the next Node
func (n *Node) Next() *Node {
	return n.next
}

// a method to return the next Node
func (n *Node) Prev() *Node {
	return n.prev
}

func main() {
	l := &List{}
	l.Push("a")
	l.Push("e")
	l.Push("i")
	l.Push("o")
	l.Push("u")

	l.Pop()

	node := l.First()
	for node != nil {
		fmt.Println(node.value)
		node = node.Next()
	}
	
	fmt.Println(l.Find("e").value)
}
