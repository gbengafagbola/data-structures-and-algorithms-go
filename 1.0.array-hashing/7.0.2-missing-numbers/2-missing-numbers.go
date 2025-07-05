package main

import "sort"

func MissingNum(nums []int) []int {
	sort.Ints(nums)

	missing := []int{}
	expected := 1

	i := 0

	for len(missing) < 2 {
		if i < len(nums) && nums[i] == expected {
			// If current number exists in nums, skip it
			for i < len(nums) && nums[i] == expected {
				i++
			}
		} else {
			// If expected is missing
			missing = append(missing, expected)
		}
		expected++
	}
	return missing
}
