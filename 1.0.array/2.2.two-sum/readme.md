### Challenge: Two Sum Problem

> Problem
given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

> Explanation

**intial thought process**
- loop through the array with pointer i,j 
- sum i & j and return if equals to target 
- continue till a pair is returned or nil 


**optimized thought process**
- Hash Map: Use a hash map to store numbers as keys and their indices as values.
- Iteration:
- - For each number in the array, calculate its complement (target - num).
- - Check if the complement is already in the hash map.
- - If found, return the indices of the complement and the current number.
- - If not found, add the current number and its index to the hash map.


```
func twoSum(nums []int, target int) []int {
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
```

**Time Complexity**
O(n): Single pass through the array with constant-time hash map operations.

**Space Complexity**
O(n): Hash map stores up to 

[source](https://leetcode.com/problems/two-sum/description/)