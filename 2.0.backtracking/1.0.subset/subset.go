package main

import (
	"fmt"
	"sort"
)

func subsets(nums []int) [][]int {
	ans := [][]int{}
	cur := []int{}
	sort.Ints(nums) // Sort the array to handle duplicates
	solution(&nums, &ans, cur, 0)
	return ans
}

func solution(nums *[]int, ans *[][]int, cur []int, index int) {
	if index > len(*nums) {
		return
	}

	// Append a copy of `cur` to `ans`
	*ans = append(*ans, append([]int(nil), cur...))

	for i := index; i < len(*nums); i++ {
		// Skip duplicates
		if i > index && (*nums)[i] == (*nums)[i-1] {
			continue
		}
		cur = append(cur, (*nums)[i])
		solution(nums, ans, cur, i+1)
		cur = cur[:len(cur)-1] // Backtrack
	}
}

func main() {
	nums := []int{1, 2, 2}
	result := subsets(nums)
	fmt.Println(result)
}