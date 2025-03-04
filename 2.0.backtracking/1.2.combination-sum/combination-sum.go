package main

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	cur := []int{}
	solution(candidates, &ans, &cur, target, 0, 0)
	return ans
}

func solution(candidates []int, ans *[][]int, cur *[]int, target int, index int, sum int) {
	if sum == target {
		*ans = append(*ans, append([]int{}, *cur...))
	} else if sum < target {
		n := len(candidates)
		for i := index; i < n; i++ {
			*cur = append(*cur, candidates[i])
			solution(candidates, ans, cur, target, i, sum+candidates[i])
			*cur = (*cur)[:len(*cur)-1]
		}
	}
}
