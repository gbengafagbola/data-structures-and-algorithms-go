package main

import "fmt"


func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j} // Return the pair
			}
		}
	}
	return nil // Return nil if no pair found
}

// Time complexity of the initial rough draft is O(n)^2

func twoSumTwo(nums []int, target int) []int {
	// Map to store numbers and their indices
	numMap := make(map[int]int)

	// Iterate through the array
	for i, num := range nums {
		complement := target - num // Calculate the complement needed for the target
		// Check if the complement exists in the map
		if index, exists := numMap[complement]; exists {
			return []int{index, i} // Return the indices of the pair
		}
		// Add the current number and its index to the map
		numMap[num] = i
	}

	return nil // Return nil if no valid pair is found
}

// Time complexity of the initial rough draft is O(n)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result) // Output: [0, 1]
}
