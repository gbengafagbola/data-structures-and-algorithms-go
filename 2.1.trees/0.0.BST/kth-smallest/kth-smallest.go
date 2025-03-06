package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Solution struct {
	res *int
	k int
}


func(s *Solution) solution(root *TreeNode) {
	if root == nil {
		return 
	}
	s.solution(root.Left)

	s.k--
	if s.k == 0 {
		s.res = &root.Val
	}
	s.solution(root.Right)
}


func kthSmallest(root *TreeNode, k int) int {
	sol := Solution{res: nil, k: k}

	sol.solution(root)
	return *sol.res
}