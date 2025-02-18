package main

import "fmt"

//mathematical approach using gauss formula
func missingNumber(nums []int) int {
	n := len(nums)
	currentSum := 0
	intendedSum := n * (n + 1) / 2

	for _, num := range nums {
		currentSum += num
	}

	return intendedSum - currentSum
}


func missingNumber2(nums []int) int {
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


func main() {
	num := []int{1,3,4,5,6}
	fmt.Println(missingNumber(num))
}