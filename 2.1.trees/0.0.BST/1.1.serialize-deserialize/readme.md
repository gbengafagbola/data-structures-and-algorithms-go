To **serialize** and **deserialize** a binary tree, we need to convert the tree into a format that can be stored or transmitted and then reconstruct it back.

A common way to do this is **preorder traversal with null markers** (to represent missing children).

---

## **Approach**
### **Serialization (Convert Tree → String)**
1. **Use preorder traversal** (Root → Left → Right).
2. **For each node**:
   - Add its value to the result.
   - If it's `nil`, add `"null"`.
3. **Join the values into a single string** using commas.

### **Deserialization (Convert String → Tree)**
1. **Split the string** into an array.
2. **Rebuild the tree recursively**:
   - Read the first value.
   - If it’s `"null"`, return `nil`.
   - Create a node and **recursively build its left and right subtrees**.

---

## **Code Implementation**
```go
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
```

---

## **How It Works**
### **Serialization**
For the tree:
```
        1
       / \
      2   3
         / \
        4   5
```
The **serialized output** would be:
```
"1,2,null,null,3,4,null,null,5,null,null"
```
- `"null"` represents missing children.
- This follows **preorder traversal** (Root → Left → Right).

---

### **Deserialization**
1. **Read values one by one**:
   - Create a node.
   - Recursively build the left subtree.
   - Recursively build the right subtree.

2. **Rebuilds the exact same tree**.

---

## **Edge Cases**
1. **Empty Tree (`nil` root)** → `"null"`
2. **Single Node Tree** → `"5,null,null"`
3. **Only Left Subtree** → `"1,2,3,null,null,null,null"`
4. **Only Right Subtree** → `"1,null,2,null,3,null,null"`

---

## **Final Thoughts**
- **Preorder traversal** ensures a structured order.
- `"null"` markers help track missing nodes.
- The solution works efficiently in **O(N) time** and **O(N) space**.

Would you like an **iterative version** using a queue? 🚀