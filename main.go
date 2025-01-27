package main

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	node *Node
}

func (t *Tree) treeInsert(value int) *Tree {
	if t.node == nil {
		t.node = &Node{value: value}
	} else {
		t.node.nodeInsert(value)
	}
	return t
}

func (n *Node) nodeInsert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.nodeInsert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.nodeInsert(value)
		}
	}
}

func (n *Node) exists(value int) bool {
	if n == nil {
		return false
	}
	if n.value == value {
		return true
	}

	println(n.value)

	if value <= n.value {
		return n.left.exists(value)
	} else {
		return n.right.exists(value)
	}
}

func printNode(n *Node) {
	if n == nil {
		return
	}

	println(n.value)
	printNode(n.left)
	printNode(n.right)
}

func main() {
	t := &Tree{}
	t.treeInsert(10).
		treeInsert(8).
		treeInsert(20).
		treeInsert(9).
		treeInsert(0).
		treeInsert(15).
		treeInsert(25)
	// printNode(t.node)

	println(t.node.exists(25))
}
