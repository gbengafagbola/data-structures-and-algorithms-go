package main

import "fmt"

type Node struct {
	Value      int
	Prev, Next *Node
}

type DoublyLinkedList struct {
	Head, Tail *Node
}

// Constructor
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// Set node as the new head
func (ll *DoublyLinkedList) SetHead(node *Node) {
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}
	ll.InsertBefore(ll.Head, node)
}

// Set node as the new tail
func (ll *DoublyLinkedList) SetTail(node *Node) {
	if ll.Tail == nil {
		ll.SetHead(node)
		return
	}
	ll.InsertAfter(ll.Tail, node)
}

// Insert nodeToInsert before node
func (ll *DoublyLinkedList) InsertBefore(node, nodeToInsert *Node) {
	if node == nil || nodeToInsert == nil || node == nodeToInsert {
		return
	}

	ll.Remove(nodeToInsert)

	nodeToInsert.Prev = node.Prev
	nodeToInsert.Next = node

	if node.Prev != nil {
		node.Prev.Next = nodeToInsert
	} else {
		ll.Head = nodeToInsert
	}

	node.Prev = nodeToInsert
}

// Insert nodeToInsert after node
func (ll *DoublyLinkedList) InsertAfter(node, nodeToInsert *Node) {
	if node == nil || nodeToInsert == nil || node == nodeToInsert {
		return
	}

	ll.Remove(nodeToInsert)

	nodeToInsert.Prev = node
	nodeToInsert.Next = node.Next

	if node.Next != nil {
		node.Next.Prev = nodeToInsert
	} else {
		ll.Tail = nodeToInsert
	}

	node.Next = nodeToInsert
}

// Insert node at a given position (1-based)
func (ll *DoublyLinkedList) InsertAtPosition(position int, nodeToInsert *Node) {
	if position <= 1 {
		ll.SetHead(nodeToInsert)
		return
	}

	currentNode := ll.Head
	currentPosition := 1

	for currentNode != nil && currentPosition != position {
		currentNode = currentNode.Next
		currentPosition++
	}

	if currentNode != nil {
		ll.InsertBefore(currentNode, nodeToInsert)
	} else {
		ll.SetTail(nodeToInsert)
	}
}

// Remove all nodes with a given value
func (ll *DoublyLinkedList) RemoveNodesWithValue(value int) {
	node := ll.Head

	for node != nil {
		nextNode := node.Next
		if node.Value == value {
			ll.Remove(node)
		}
		node = nextNode
	}
}

// Remove a specific node
func (ll *DoublyLinkedList) Remove(node *Node) {
	if node == nil {
		return
	}

	if node == ll.Head {
		ll.Head = node.Next
	}
	if node == ll.Tail {
		ll.Tail = node.Prev
	}

	ll.removeBinding(node)
}

// Check if list contains a node with a given value
func (ll *DoublyLinkedList) ContainsNodeWithValue(value int) bool {
	node := ll.Head
	for node != nil {
		if node.Value == value {
			return true
		}
		node = node.Next
	}
	return false
}

// Unlink node from list
func (ll *DoublyLinkedList) removeBinding(node *Node) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	node.Prev = nil
	node.Next = nil
}

// Helper to print list values (forward)
func (ll *DoublyLinkedList) PrintForward() {
	node := ll.Head
	for node != nil {
		fmt.Printf("%d ", node.Value)
		node = node.Next
	}
	fmt.Println()
}

func main() {
	ll := NewDoublyLinkedList()

	a := &Node{Value: 1}
	b := &Node{Value: 2}
	c := &Node{Value: 3}
	d := &Node{Value: 4}

	ll.SetHead(a)
	ll.InsertAfter(a, b)
	ll.InsertAfter(b, c)
	ll.InsertAtPosition(2, d) // Inserts d between a and b

	ll.PrintForward() // Output: 1 4 2 3
}
