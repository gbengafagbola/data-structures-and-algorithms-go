package main

type TreeNode struct {
	Val int
	Right *TreeNode
	Left *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{root}
	ans := [][]int{}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, level)
	}
	return ans
}