package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func zigzag(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var ans [][]int

	q := []*TreeNode{root}
	levelIndex := 1

	for len(q) > 0 {
		level := make([]int, 0)
		qLen := len(q)
		for i := 0; i < qLen; i++ {
			current := q[0]
			q = q[1:]

			level = append(level, current.Val)

			if current.Left != nil {
				q = append(q, current.Left)
			}

			if current.Right != nil {
				q = append(q, current.Right)
			}
		}

		if levelIndex%2 == 0 {
			for i, j := 0, len(level)-1; i < j; i, j = i+1, j-1 {
				level[i], level[j] = level[j], level[i]
			}
		}
		ans = append(ans, level)
		levelIndex++
	}
	return ans
}
