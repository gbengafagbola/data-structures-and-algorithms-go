package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Serialize: Convert tree to a string
func serialize(root *TreeNode) string {
	var sb []string
	serializeHelper(root, &sb)
	return strings.Join(sb, ",")
}

// Helper function for serialization (Preorder Traversal)
func serializeHelper(node *TreeNode, sb *[]string) {
	if node == nil {
		*sb = append(*sb, "null")
		return
	}
	*sb = append(*sb, fmt.Sprintf("%d", node.Val)) // Add node value
	serializeHelper(node.Left, sb)                 // Serialize left subtree
	serializeHelper(node.Right, sb)                // Serialize right subtree
}

// Deserialize: Convert string back to tree
func deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	return deserializeHelper(&nodes)
}

// Helper function for deserialization
func deserializeHelper(nodes *[]string) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}

	val := (*nodes)[0]
	*nodes = (*nodes)[1:] // Move to next element

	if val == "null" {
		return nil
	}

	// Convert string to int
	var num int
	fmt.Sscanf(val, "%d", &num)
	node := &TreeNode{Val: num}

	node.Left = deserializeHelper(nodes)  // Rebuild left subtree
	node.Right = deserializeHelper(nodes) // Rebuild right subtree

	return node
}

// Example Usage
func main() {
	// Example tree:
	//        1
	//       / \
	//      2   3
	//         / \
	//        4   5
	root := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}}

	// Serialize
	data := serialize(root)
	fmt.Println("Serialized:", data)

	// Deserialize
	newRoot := deserialize(data)
	fmt.Println("Deserialized Root:", newRoot.Val) // Should be 1
}
