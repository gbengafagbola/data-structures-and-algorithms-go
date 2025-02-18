package main

func missingNumber(nums []int) int {

	n := len(nums)

	currentSum := 0
	intendedSum := 0

	for i := 0; i < n; i++ {
		currentSum += i
	}

	for i := 0; i < (n+1); i++ {
		intendedSum += i
	}

	return intendedSum - currentSum
}