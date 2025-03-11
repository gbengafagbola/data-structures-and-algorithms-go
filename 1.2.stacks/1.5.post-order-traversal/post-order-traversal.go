package main

type TreeNode struct {
	Val int
	Right *TreeNode
	Left *TreeNode
}

func postOrderTraversal(root *TreeNode) []int {
	var ans []int

	if root == nil {
		return ans
	}

	var s1 []*TreeNode
	var s2 []*TreeNode


	s1 = append(s1, root)

	for len(s1) > 0 {
		x := s1[len(s1) - 1]
		s1 = s1[:len(s1) - 1]

		if x.Left != nil {
			s1 = append(s1, x.Left)
		}
		if x.Right != nil {
			s1 = append(s1, x.Right)
		}
	}

	for len(s2) > 0 {
		y := s2[len(s2) - 1]
		s2 = s2[:len(s2) - 1]
		ans = append(ans, y.Val)
	}
	return ans
}