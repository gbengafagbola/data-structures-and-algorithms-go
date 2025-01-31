package main

import "fmt"

func BinarySearch (data []int, value int) bool {
	var mid int
	// the first character in the array
	start := 0
	// the last character in the array 
	end := len(data) - 1
	// provided as long as 0 is not greater than end
	for start <= end {
		// median value 
		mid = (start + end)/2

		if data[mid] == value {
			return true // if value is found return true
		} else {
			if data[mid] < value {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return false // value not found
}


// Testing

func main() {
	arr := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println("Binary Search:", BinarySearch(arr, 8))
}


//  Time Complexity
// O(log(n))