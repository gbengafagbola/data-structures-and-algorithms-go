package main

// Do not edit the class below except
// for the depthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	// Write your code here.
    array = append(array, n.Name)
    
    for _, child := range n.Children {
      array = child.DepthFirstSearch(array)
    }
    
	return array
}
