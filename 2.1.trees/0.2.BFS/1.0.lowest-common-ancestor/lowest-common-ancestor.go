package main

type TreeNode struct {
	Val int
	Right *TreeNode
	Left *TreeNode
}

func lowestCommonAncestor(node, p, q *TreeNode) *TreeNode {

	if node == nil {
		return nil
	}

	if node == p || node == q {
		return node
	}

	left := lowestCommonAncestor(node.Left, p, q)
	right := lowestCommonAncestor(node.Right, p, q)

	if left != nil && right != nil {
		return node
	}

	if left != nil {
		return left
	}

	return right
}