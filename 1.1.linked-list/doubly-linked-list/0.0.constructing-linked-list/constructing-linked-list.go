package main

type Node struct {
	Value      int
	Prev, Next *Node
}

type DoublyLinkedList struct {
	Head, Tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {

	return nil
}

func (ll *DoublyLinkedList) SetHead(node *Node) {

}

func (ll *DoublyLinkedList) SetTail(node *Node) {

}

func (ll *DoublyLinkedList) InsertBefore(node, nodeToInsert *Node) {

    // node.Prev.Next = nodeToInsert
	if node == ll.Head  && node == ll.Tail {
		return
	}

	// if node exists already in the ll
	ll.Remove(nodeToInsert)
    node.Prev = nodeToInsert.Prev
	nodeToInsert.Next = node
	ll.removeBinding(node)
	if node.Prev == nil { 
		ll.Head = nodeToInsert
	} else {
		node.Prev.Next = nodeToInsert
	}
	node.Prev = nodeToInsert
}

func (ll *DoublyLinkedList) InsertAfter(node, nodeToInsert *Node) {
	if node == ll.Head && node == ll.Tail {
		return
		// ll.Tail = nodeToInsert
		// node.Next = nodeToInsert
	}

}

func (ll *DoublyLinkedList) InsertAtPosition(position int, nodeToInsert *Node) {
	 
}

func (ll *DoublyLinkedList) RemoveNodesWithValue(Value int) {
	node := ll.Head

	for node != nil {
		NextNode := node.Next // Save Next before possibly removing current

		if node.Value == Value {
			ll.Remove(node)
		}

		node = NextNode
	}
}


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


func (ll *DoublyLinkedList) ContainsNodeWithValue(Value int) bool {
	node := ll.Head

	for node != nil {
		if node.Value == Value {
			return true
		}
		node = node.Next
	}

	return false
}


// Internal helper to unlink a node from its neighbors
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