package main

// Node represents a single element in the binary search tree.
// Each node contains a value, a pointer to the left child, and a pointer to the right child.
type Node struct {
	value int
	left  *Node
	right *Node
}

// Tree represents the binary search tree.
// It contains a pointer to the root node.
type Tree struct {
	node *Node
}

// treeInsert inserts a value into the tree.
// If the tree is empty, it creates the root node; otherwise, it delegates to nodeInsert.
func (t *Tree) treeInsert(value int) *Tree {
	if t.node == nil {
		// Create the root node if the tree is empty.
		t.node = &Node{value: value}
	} else {
		// Delegate the insertion to the nodeInsert method of the root node.
		t.node.nodeInsert(value)
	}
	return t // Return the tree to allow method chaining.
}

// nodeInsert inserts a value into the binary search tree starting from the current node.
func (n *Node) nodeInsert(value int) {
	if value <= n.value {
		// If the value is less than or equal to the current node's value,
		// it should go to the left subtree.
		if n.left == nil {
			// If the left child is nil, create a new node and assign it.
			n.left = &Node{value: value}
		} else {
			// If the left child exists, recursively call nodeInsert on the left child.
			n.left.nodeInsert(value)
		}
	} else {
		// If the value is greater than the current node's value,
		// it should go to the right subtree.
		if n.right == nil {
			// If the right child is nil, create a new node and assign it.
			n.right = &Node{value: value}
		} else {
			// If the right child exists, recursively call nodeInsert on the right child.
			n.right.nodeInsert(value)
		}
	}
}

// exists checks whether a value exists in the binary search tree starting from the current node.
func (n *Node) exists(value int) bool {
	if n == nil {
		// If the node is nil, the value does not exist.
		return false
	}
	if n.value == value {
		// If the current node's value matches the search value, return true.
		return true
	}

	// Print the current node's value during traversal (for debugging/visualization).
	println(n.value)

	if value <= n.value {
		// If the search value is less than or equal to the current node's value,
		// search in the left subtree.
		return n.left.exists(value)
	} else {
		// If the search value is greater, search in the right subtree.
		return n.right.exists(value)
	}
}

// printNode performs a pre-order traversal of the tree,
// printing the value of each node as it is visited.
func printNode(n *Node) {
	if n == nil {
		// If the node is nil, there's nothing to print; return.
		return
	}

	// Print the current node's value.
	println(n.value)

	// Recursively print the left and right subtrees.
	printNode(n.left)
	printNode(n.right)
}

func main() {
	// Create an empty binary search tree.
	t := &Tree{}

	// Insert values into the tree using method chaining.
	t.treeInsert(10). // Root node
		treeInsert(8).  // Left child of root
		treeInsert(20). // Right child of root
		treeInsert(9).  // Right child of node 8
		treeInsert(0).  // Left child of node 8
		treeInsert(15). // Left child of node 20
		treeInsert(25)  // Right child of node 20

	// Uncomment the following line to print the entire tree structure (pre-order traversal).
	printNode(t.node)

	// Check if specific values exist in the tree and print the result.
	println(t.node.exists(25)) // Outputs: true (Value exists in the tree)
}


///[10 8 0 9 20 15 25]
//	 	10
//	   /  \
//	 8    20
//  / \   /  \
// 0   9 15   25