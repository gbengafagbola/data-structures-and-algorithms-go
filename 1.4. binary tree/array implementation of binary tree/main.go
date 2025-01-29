package main

import "fmt"

type Tree struct {
	nodes []int
}

// Insert a value into the tree.
func (t *Tree) Insert(value int) {
	t.nodes = append(t.nodes, value)
}

// Check if a value exists in the tree.
func (t *Tree) Exists(value int) bool {
	for _, v := range t.nodes {
		if v == value {
			return true
		}
	}
	return false
}

// Print the tree (level-order traversal).
func (t *Tree) PrintTree() {
	for i, v := range t.nodes {
		fmt.Printf("Index %d: Value %d\n", i, v)
	}
}

func main() {
	t := &Tree{}

	// Insert values into the tree.
	t.Insert(10) // Root
	t.Insert(8)  // Left child of root
	t.Insert(20) // Right child of root
	t.Insert(9)  // Left child of node at index 1
	t.Insert(0)  // Right child of node at index 1
	t.Insert(15) // Left child of node at index 2
	t.Insert(25) // Right child of node at index 2

	// Print the tree.
	fmt.Println("Tree structure:")
	t.PrintTree()

	// Check if values exist in the tree.
	fmt.Println("\nValue 25 exists:", t.Exists(25)) // true
	fmt.Println("Value 30 exists:", t.Exists(30))   // false
}
