package main

import "fmt"

func moveZeroes(nums []int) []int {
	zeroIndex := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != zeroIndex {
				nums[i], nums[zeroIndex] = nums[zeroIndex], nums[i]
			}
			zeroIndex++
		}
	}
	return nums
}


func moveZeroes2(nums []int) []int {
	zeroIndex := 0
	n := len(nums)

	for nonZeroIndex := 0; nonZeroIndex < n; nonZeroIndex++ {
		if nums[nonZeroIndex] != 0 {
			nums[zeroIndex] = nums[nonZeroIndex]
			zeroIndex++
		}
	}
	for i := zeroIndex; i < n; i++ {
		nums[i] = 0
	}

	return nums
}

func main() {
	nums := []int{2, 1, 0, 4, 0, 0, 6}
	result := moveZeroes(nums)
	fmt.Println(result)
}
