package structural

import (
	"testing"
)

// TestBinaryTree tests the insertion and in-order traversal of the binary tree.
func TestBinaryTree(t *testing.T) {
	// Create root node
	root := &Node{Value: 10}

	// Insert elements
	root.Insert(5)
	root.Insert(15)
	root.Insert(2)
	root.Insert(7)
	root.Insert(12)
	root.Insert(20)

	// Expected sorted order from in-order traversal: [2, 5, 7, 10, 12, 15, 20]
	expected := []int{2, 5, 7, 10, 12, 15, 20}
	var result []int

	// Helper function to capture traversal output
	var captureTraversal func(*Node)
	captureTraversal = func(n *Node) {
		if n == nil {
			return
		}
		captureTraversal(n.Left)
		result = append(result, n.Value)
		captureTraversal(n.Right)
	}

	// Perform in-order traversal and capture results
	captureTraversal(root)

	// Compare expected and actual results
	if len(result) != len(expected) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}

	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %v at index %d but got %v", v, i, result[i])
		}
	}
}
