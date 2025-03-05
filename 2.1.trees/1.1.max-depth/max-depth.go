package main


type TreeNode struct {
	Val	int
	Right *TreeNode
	Left *TreeNode
}

func maxDepth (root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := 1 + maxDepth(root.Left)
	right := 1 + maxDepth(root.Right)
	ans := max(left, right)
	return ans
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}