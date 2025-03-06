package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func hasSum(root *TreeNode, sum int, cur int) bool {
	cur += root.Val

	if cur == sum && root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil {
		if hasSum(root.Left, sum, cur) {
			return true
		}
	}

	if root.Right != nil {
		if hasSum(root.Right, sum, cur) {
			return true
		}
	}

	return false
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	return hasSum(root, sum, 0)
}
