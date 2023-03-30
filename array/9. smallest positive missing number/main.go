package main

import "fmt"

func SmallestPositiveMissingNumber(arr []int, size int) int {
	i := 0
	missing := -1

	for i < size {
		currentIndex := i + 1

		if arr[i] != currentIndex && arr[i] > size {
			missing = currentIndex
			i++
		}

		if currentIndex < size && arr[i] != currentIndex {
			arr[i], arr[arr[i]-1] = arr[arr[i]-1], arr[i]
		}

		if currentIndex < size && arr[i] == i+1 {
			i++
		}
	}

	return missing
}

func main() {
	array := []int{1, 3, 4, 2, 6}
	size := len(array)
	fmt.Println("array", SmallestPositiveMissingNumber(array, size))
}

//  Time Complexity
// O(n)
