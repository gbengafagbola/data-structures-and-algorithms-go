package main

import "fmt"

func moveZeroes(nums []int) []int {
	aZeroIndex := 0

	for i := 0; i < len(nums); i++ {
		//we only have intrest with non-zero values in the nums array 
		if nums[i] != 0 {  
			//we have to ensure that the ith index is also not the same as the zeroIndex
			if i != aZeroIndex {
				nums[i], nums[aZeroIndex] = nums[aZeroIndex], nums[i]
			}
			aZeroIndex++
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
