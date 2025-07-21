package main

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

// Insert inserts a value into the BST Avg time complexity O(log n) Time O(1) space WOrst O(n) time O(1) space
func (tree *BST) Insert(value int) *BST {
	currentNode := tree

	for {
		if value < currentNode.Value {
			if currentNode.Left == nil {
				currentNode.Left = &BST{Value: value}
				break
			}
			currentNode = currentNode.Left
		} else {
			if currentNode.Right == nil {
				currentNode.Right = &BST{Value: value}
				break
			}
			currentNode = currentNode.Right
		}
	}
	return tree
}

// Contains checks if a value exists in the BST
func (tree *BST) Contains(value int) bool {
	currentNode := tree

	for currentNode != nil {
		if value < currentNode.Value {
			currentNode = currentNode.Left
		} else if value > currentNode.Value {
			currentNode = currentNode.Right
		} else {
			return true
		}
	}
	return false
}

// Remove deletes a value from the BST
func (tree *BST) Remove(value int) *BST {
	return removeNode(tree, value, nil)
}

func removeNode(current *BST, value int, parent *BST) *BST {
	node := current

	for node != nil {
		if value < node.Value {
			parent = node
			node = node.Left
		} else if value > node.Value {
			parent = node
			node = node.Right
		} else {
			// Case 1: Node has two children
			if node.Left != nil && node.Right != nil {
				// Replace with smallest value in right subtree
				node.Value = getMinValue(node.Right)
				// Remove the duplicate from right subtree
				node.Right = removeNode(node.Right, node.Value, node)
			} else if parent == nil {
				// Removing root node with 1 or 0 children
				if node.Left != nil {
					*node = *node.Left
				} else if node.Right != nil {
					*node = *node.Right
				} else {
					// Single root node
					node = nil
				} 
			} else if parent.Left == node {
				if node.Left != nil {
					parent.Left = node.Left
				} else {
					parent.Left = node.Right
				}
			} else if parent.Right == node {
				if node.Left != nil {
					parent.Right = node.Left
				} else {
					parent.Right = node.Right
				}
			}
			break
		}
	}
	return current
}

func getMinValue(node *BST) int {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current.Value
}
