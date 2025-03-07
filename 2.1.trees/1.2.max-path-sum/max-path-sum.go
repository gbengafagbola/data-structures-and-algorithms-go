package main

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


var ans int


func solution(node *TreeNode) int{
	if node == nil {
		return 0
	}

	left := solution(node.Left)
	right := solution(node.Right)

	maxSide := int(math.Max(float64(node.Val), float64(node.Val + int(math.Max(float64(left), float64(right))))))
	maxCurrent := int(math.Max(float64(maxSide), float64(node.Val + left + right)))

	ans = int(math.Max(float64(ans), float64(maxCurrent)))

	return maxSide
}


func maxPathSum(root *TreeNode) int {
	ans = math.MinInt32
	solution(root)
	return ans
}