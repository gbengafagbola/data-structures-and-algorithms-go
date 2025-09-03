package main

type Node struct {
	Name     string
	Children []*Node
}

// BreadthFirstSearch performs BFS and returns the traversal order.
func (n *Node) BreadthFirstSearch(array []string) []string {
	// Initialize queue with the root node
	queue := []*Node{n}

	for len(queue) > 0 {
		// Pop the first node in the queue
		currentNode := queue[0]
		queue = queue[1:]

		// Add the node's name to the result array
		array = append(array, currentNode.Name)

		// Add all children of the current node to the queue
		queue = append(queue, currentNode.Children...)
	}

	return array
}
