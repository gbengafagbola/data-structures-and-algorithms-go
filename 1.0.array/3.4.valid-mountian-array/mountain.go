package main

import "fmt"

func validMountainArray(arr []int) bool {
	index := 1

	for index < len(arr) && arr[index] > arr[index-1] {
		index++
	}

	if index == 1 || index == len(arr) {
		return false
	}

	for index < len(arr) && arr[index] <= arr[index-1] {
		index++
	}
	return index == len(arr)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}

	result := validMountainArray(arr)

	fmt.Print(result)
}
