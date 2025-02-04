package structural

import "fmt"

// Node represents a single node in a binary tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert adds a new node to the binary tree
func (n *Node) Insert(value int) {
	if value <= n.Value {
		// Insert into the left subtree
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		// Insert into the right subtree
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

// InOrderTraversal prints the tree values in ascending order
func (n *Node) InOrderTraversal() {
	if n == nil {
		return
	}
	n.Left.InOrderTraversal()
	fmt.Print(n.Value, " ")
	n.Right.InOrderTraversal()
}

func main() {
	// Create root node
	root := &Node{Value: 10}

	// Insert elements
	root.Insert(5)
	root.Insert(15)
	root.Insert(2)
	root.Insert(7)
	root.Insert(12)
	root.Insert(20)

	// Print sorted tree elements
	fmt.Print("InOrder Traversal: ")
	root.InOrderTraversal()
	fmt.Println()
}
