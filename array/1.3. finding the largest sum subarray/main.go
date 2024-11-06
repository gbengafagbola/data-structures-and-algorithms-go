package main

import "fmt"

func MaxSubArraySum(data []int) int {
	sum := 0
	max := 0
	length := len(data) - 1

	for i := 0; i <= length; i++ {
		sum = sum + data[i]

		if sum < 0 {
			sum = 0
		}

		if max < sum {
			max = sum
		}
	}
	return max
}

// Naive solution of finding maximum sub array sum
func NaiveMaxSubArraySum(data []int) int {
	max := data[0]
	sum := 0
	length := len(data) - 1

	for i := 0; i <= length; i++ {
		for j := i; j <= length; j++ {
			sum += data[j]
			if sum > max {
				max = sum
			}
		}
		sum = 0
	}
	return max
}

// Testing
func main() {
	arr := []int{1, -2, 3, 4, -4, 6, -14}
	fmt.Println("MaxSubArraySum:", MaxSubArraySum(arr))
	// fmt.Println("MaxSubArraySum:", NaiveMaxSubArraySum(arr))
}

//  Time Complexity MaxSubArraySum
// O(n)
//  Time Complexity NaiveMaxSubArraySum
// O(n^2)

