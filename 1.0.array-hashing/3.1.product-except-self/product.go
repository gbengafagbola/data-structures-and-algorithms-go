package main

import "fmt"





func productExceptSelf(nums []int) []int {
	result := []int{}

	track := 0
	multiple := 1

	var i int = 0;

	for i < len(nums) && track <= len(nums){
		if track != i {
			multiple = multiple * nums[i]
		}
		i++
	}
	result = append(result, multiple)
	i = 0;
	track++

	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf(nums)
	fmt.Println(result) // Output: [24, 12, 8, 6]
}



func productExceptSelfX(nums []int) []int {
	result := []int{}
	track := 0

	for track < len(nums) {
		multiple := 1
		for i := 0; i < len(nums); i++ {
			if track != i {
				multiple *= nums[i]
			}
		}
		result = append(result, multiple)
		track++
	}

	return result
}