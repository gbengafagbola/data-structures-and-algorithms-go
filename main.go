package main

import (
	"fmt"
)

type List struct {
	head   *Node
	tail   *Node
	length int
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Push(value int) {
	node := &Node{value: value}

	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
	l.length = +1
}

func (l *List) Pop() {
	l.tail = l.tail.prev
	l.tail.next = nil
}

func (l *List) Find(value int) (*Node) {
	node := l.First()

	for node.next != nil {
		if (node.value == value) {
			return node
		}
		node = node.next
	}

	return &Node{value: 0}
}

type Node struct {
	value int
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
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)

	l.Pop()

	node := l.First()
	for node != nil {
		fmt.Println(node.value)
		node = node.Next()
	}
	
	fmt.Println(l.Find(2).value)
}
