
package main

// import "fmt"

// H                       T
// A --> E --> I --> O --> U
// e     i     o     u

type Node struct {
	value string
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) Push (value string) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node 
	} else {
		l.tail.next = node
	}
	l.tail = node
}

func (l *List) First() *Node {
	return l.head;
}

func (l *List) Last() *Node {
	return l.tail;
}

func (n *Node) Next() *Node {
	return n.next
}

func (l *List) Pop() {
	// If the list is empty, do nothing
	if l.head == nil {
		return
	}

	// If the list has only one node
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return
	}

	// Traverse to the second-to-last node
	pointer := l.head
	for pointer.next != l.tail {
		pointer = pointer.next
	}

	// Update the tail and remove the last node
	pointer.next = nil
	l.tail = pointer
}


func main(){
  list := List{}
   list.Push("a")
   list.Push("e")
   list.Push("i")
   list.Push("o")
   list.Push("u")

//    node := list.First()
//    fmt.Println(node.value)

// node := list.First()
// fmt.Println(node.Next().value)

//  n := list.First()
//  for {
// 	fmt.Println(n.value)
// 	n = n.Next()
// 	if n == nil {
// 		break
// 	}
//  }

// list.Pop()
}