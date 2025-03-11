package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postOrderTraversal(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}

	var s1, s2 []*TreeNode

	// Step 1: Push root to stack s1
	s1 = append(s1, root)

	for len(s1) > 0 {
		// Pop from s1
		x := s1[len(s1)-1]
		s1 = s1[:len(s1)-1]

		// Push to s2 (to reverse the order)
		s2 = append(s2, x)

		// Push left first, then right (reversed post-order)
		if x.Left != nil {
			s1 = append(s1, x.Left)
		}
		if x.Right != nil {
			s1 = append(s1, x.Right)
		}
	}

	// Step 2: Pop from s2 to get post-order traversal
	for len(s2) > 0 {
		y := s2[len(s2)-1]
		s2 = s2[:len(s2)-1]
		ans = append(ans, y.Val)
	}

	return ans
}

func main() {
	// Creating a test tree
	root := &TreeNode{1,
		&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}},
		&TreeNode{3, nil, &TreeNode{6, nil, nil}},
	}

	fmt.Println(postOrderTraversal(root)) // Output: [4, 5, 2, 6, 3, 1]
}
