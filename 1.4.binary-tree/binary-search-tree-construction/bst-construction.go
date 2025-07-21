package main

import "fmt" // Added for example usage

// BST represents a Binary Search Tree node.
type BST struct {
	Value int
	Left  *BST
	Right *BST
}

// NewBST creates and returns a new BST with the given value.
func NewBST(value int) *BST {
	return &BST{Value: value}
}

// --- Insert Operation
// Insert inserts a value into the BST.
// Average: O(log n) time | O(1) space
// Worst: O(n) time | O(1) space (for a skewed tree)
func (tree *BST) Insert(value int) *BST {
	currentNode := tree
	for {
		if value < currentNode.Value {
			if currentNode.Left == nil {
				currentNode.Left = &BST{Value: value}
				break // Value inserted, exit loop
			}
			currentNode = currentNode.Left
		} else { // value >= currentNode.Value (handles duplicates by placing them on the right)
			if currentNode.Right == nil {
				currentNode.Right = &BST{Value: value}
				break // Value inserted, exit loop
			}
			currentNode = currentNode.Right
		}
	}
	return tree // Return the original tree root
}



// Contains checks if a value exists in the BST.
// Average: O(log n) time | O(1) space
// Worst: O(n) time | O(1) space (for a skewed tree)
func (tree *BST) Contains(value int) bool {
	currentNode := tree
	for currentNode != nil {
		if value < currentNode.Value {
			currentNode = currentNode.Left
		} else if value > currentNode.Value {
			currentNode = currentNode.Right
		} else {
			return true // Value found
		}
	}
	return false // Value not found after traversing
}




// Remove deletes a value from the BST.
// Average: O(log n) time | O(1) space
// Worst: O(n) time | O(1) space (for a skewed tree)
func (tree *BST) Remove(value int) *BST {
	// The helper function removeNode handles the actual removal logic.
	// We pass nil as the parent for the initial call, as the root has no parent.
	return removeNode(tree, value, nil)
}

// removeNode is a recursive helper function to handle the actual node removal logic.
// It returns the new subtree root after removal.
func removeNode(node *BST, value int, parent *BST) *BST {
	if node == nil {
		return nil // Value not found, or empty tree/subtree
	}

	if value < node.Value {
		// Traverse left, updating the left child of the current node
		node.Left = removeNode(node.Left, value, node)
	} else if value > node.Value {
		// Traverse right, updating the right child of the current node
		node.Right = removeNode(node.Right, value, node)
	} else {
		// Value found! Now handle the removal cases.

		// Case 1: Node has two children
		if node.Left != nil && node.Right != nil {
			// Find the smallest value in the right subtree (in-order successor)
			// This value will replace the current node's value.
			node.Value = getMinValue(node.Right)
			// Recursively remove the duplicate (the smallest value) from the right subtree.
			node.Right = removeNode(node.Right, node.Value, node)
		} else {
			// Case 2: Node has one or zero children

			// Determine the child to replace the current node.
			// If node.Left is not nil, it's the replacement. Otherwise, node.Right is (could be nil).
			var childToReplace *BST
			if node.Left != nil {
				childToReplace = node.Left
			} else {
				childToReplace = node.Right
			}

			// If the node being removed is the root of the entire tree (parent is nil)
			if parent == nil {
				// Special handling for root: the new root becomes its child (or nil if no children).
				return childToReplace
			}

			// If not the root, update the parent's child pointer
			if parent.Left == node {
				parent.Left = childToReplace
			} else { // parent.Right == node
				parent.Right = childToReplace
			}
			return node // Return the original node (now linked out of the tree)
		}
	}
	return node // Return the current node (or its updated subtree)
}

// getMinValue finds the smallest value in a given subtree.
// It traverses the left child until it finds a node with no left child.
func getMinValue(node *BST) int {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current.Value
}


func main() {
	// Create a new BST
	tree := NewBST(10)

	// Insert values
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(13)
	tree.Insert(22)
	tree.Insert(1)
	tree.Insert(6)
	tree.Insert(8)
	tree.Insert(12)
	tree.Insert(14)
	tree.Insert(20)
	tree.Insert(25)

	fmt.Println("Initial BST (in-order traversal, conceptually):")
	// You would typically implement an in-order traversal to see the sorted values
	// For simplicity here, just showing operations.
	fmt.Printf("Tree root value: %d\n", tree.Value)
	fmt.Printf("Does 7 exist? %v\n", tree.Contains(7))    // true
	fmt.Printf("Does 100 exist? %v\n", tree.Contains(100)) // false

	fmt.Println("\n--- Removing Nodes ---")

	// Test removing a leaf node
	fmt.Println("Removing 1 (leaf node)...")
	tree.Remove(1)
	fmt.Printf("Does 1 exist? %v\n", tree.Contains(1)) // false

	// Test removing a node with one child
	fmt.Println("Removing 2 (node with one child)...")
	tree.Remove(2)
	fmt.Printf("Does 2 exist? %v\n", tree.Contains(2)) // false
	fmt.Printf("Does 6 exist? %v\n", tree.Contains(6)) // true (should still be there)

	// Test removing a node with two children
	fmt.Println("Removing 15 (node with two children)...")
	tree.Remove(15)
	fmt.Printf("Does 15 exist? %v\n", tree.Contains(15)) // false
	// The value 20 should have replaced 15 (smallest in right subtree)
	fmt.Printf("Does 20 exist? %v\n", tree.Contains(20)) // true

	// Test removing the root
	fmt.Println("Removing 10 (root node)...")
	newRoot := tree.Remove(10) // Root might change, so capture the new root
	if newRoot != nil {
		fmt.Printf("New root value: %d\n", newRoot.Value) // Should be 12 (original 10's successor)
	} else {
		fmt.Println("Tree is now empty.")
	}
	fmt.Printf("Does 10 exist? %v\n", newRoot.Contains(10)) // false
	fmt.Printf("Does 12 exist? %v\n", newRoot.Contains(12)) // true
}