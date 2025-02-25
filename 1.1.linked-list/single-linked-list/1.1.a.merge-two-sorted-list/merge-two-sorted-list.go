package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type List struct {
	head *Node
	tail *Node
}

// Merges two sorted arrays into a linked list
func (l *List) mergeSort(arr1, arr2 []int) []int {
	l1, l2 := 0, 0

	for l1 < len(arr1) && l2 < len(arr2) {
		Node1 := &Node{value: arr1[l1]}
		Node2 := &Node{value: arr2[l2]}

		if Node1.value < Node2.value {
			l.append(Node1)
			l1++
		} else {
			l.append(Node2)
			l2++
		}
	}

	// Add remaining elements from arr1
	for l1 < len(arr1) {
		l.append(&Node{value: arr1[l1]})
		l1++
	}

	// Add remaining elements from arr2
	for l2 < len(arr2) {
		l.append(&Node{value: arr2[l2]})
		l2++
	}

	return l.toSlice()
}

// Appends a node to the linked list
func (l *List) append(node *Node) {
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node // Move the tail forward
	}
}

// Converts the linked list back to a slice
func (l *List) toSlice() []int {
	var result []int
	current := l.head
	for current != nil {
		result = append(result, current.value)
		current = current.next
	}
	return result
}

func main() {
	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{1, 2, 4, 6}

	list := &List{} // Create an instance of List
	result := list.mergeSort(arr1, arr2)

	fmt.Println(result) // Output: [1, 1, 2, 3, 4, 5, 6, 7]
}
