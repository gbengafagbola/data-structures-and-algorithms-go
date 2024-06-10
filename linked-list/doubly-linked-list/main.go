package main

// import "fmt"

// defining a List of type structure with a field called head which points to a Node that is the first element in the List.
// and tail a which points to a Node that is the last element in the List

type List struct {
	head   *Node
	tail   *Node
	length int
}

// the first method / func would give the first element of the list or rather return the head Node in the List
func (l *List) First() *Node {
	return l.head
}

// the last method / func would give the last element of the list or rather return the tail Node in the List
func (l *List) Last() *Node {
	return l.tail
}

// the second method / func here push is to add an item to the List
func (l *List) Push(value int) {
	//initializing node from passed in value
	node := &Node{value: value}

	// if the list has no head element, make the node the head of the list
	// else: from the first iteration where there was no head, we made the passed in value the head and also the tail but in the second iteration of push we would make the incoming value sit after the current tail and head (which are of cause of same node) by calling the next method off the tail that is l.tail.next = node
	// since it's a doubly linked list, we have to keep track of both the next and previous element of each node, hence assigning the previous tail of the list to the new nodes previous element.
	// and finally we have a new tail element which is the new node
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
	// A --> B --> C --> D
	// idea 1
	// get the second to last node and remove the next refrence
	// that is to pop off D, traverse (get access) to C and remove its refrence to D
		// Sub Idea to get to C
		// idea 1
		// traverse the list via length-1

	stopper := l.length - 1

	for i := 0; i <= stopper; i++ {
		n := l.head
		if i == stopper {
			n.next = nil
			l.tail = n
			l.length--
		}
		n = n.next

	}

	// return l.tail
	// l.tail
	// start.next
}

// func (l *List) Find(value int) *Node {
// 	node := &Node{value: value}

// 	for e := l.First(); e != nil; e = e.Next() {
// 		if e == node {
// 			return node
// 		}
// 	}

// 	return node
// }

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

	n := l.First()
	for {
		println(n.value)
		n = n.Next()
		if n == nil {
			break
		}
	}


	println(l)
	l.Pop()

	// l.Find(2)



	// n := l.Last()
	// for {
	// 	println(n.value)
	// 	n = n.Prev()
	// 	if n == nil {
	// 		break
	// 	}
	// }

}
