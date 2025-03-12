package main

import "math"

func rob(nums []int) int {
	n := len(nums)
	if len(nums) == 1 {
		return nums[0]
	}

	house1 := nums[0]
	house2 := int(math.Max(float64(nums[0]), float64(nums[1])))

	ans := house2
	for i := 2; i < n; i++ {
		ans = int(math.Max(float64(house2), float64(nums[i])+float64(house1)))
		house1 = house2
		house2 = ans
	}
	return ans
}