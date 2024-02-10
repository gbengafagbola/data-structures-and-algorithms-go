package main

// import "fmt"

// defining a List of type structure with a field called head which points to a Node that is the first element in the List.
// and tail a which points to a Node that is the last element in the List
type List struct {
	head *Node
	tail *Node
}

// the first method / func would give the first element of the list or rather return the first Node in the List 
func (l *List) First() *Node {
	return l.head;
} 


// the second method / func here push is to add an item to the List 
func (l *List) Push(value int) {
	//initializing node from passed in value
	node := &Node{value: value};

	// if the list has no head element, make the node the head of the list
	// else: from the first iteration where there was no head, we made the passed in value the head and also the tail but in the second iteration of push we would make the incoming value sit after the current tail and head (which are of cause of same node) by calling the next method off the tail that is l.tail.next = node
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
}


type Node struct {
	value int
	next *Node
}
// a method to return the next Node
func (n *Node) Next() *Node{
	return n.next
}



func main() {
	l := &List{};
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

}
