package main

import (
	"fmt"
	"math"
)

func targetFirstLast(arr []int, target int) []int {
	if len(arr) == 0 {
		return []int{0, 0}
	}

	left := 0
	right := len(arr) - 1

	for left <= right {

		if arr[left] != target {
			left++
		}

		if arr[right] != target {
			right--
		}

		if arr[left] == target && arr[right] == target {
			return []int{left, right}
		}
	}
	return []int{-1, -1}
}

func targetFirstLast2(arr []int, target int) []int {
	first := findFirst(arr, target)
	last := findLast(arr, target)
	return []int{first, last}
}

func findFirst(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := int(math.Floor((float64(left) + float64(right)) / 2))
		if arr[mid] == target {
			if mid == 0 || arr[mid-1] != target {
				return mid
			} else {
				right = mid - 1
			}
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func findLast(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := int(math.Floor((float64(left) + float64(right)) / 2))
		if arr[mid] == target {
			if mid == len(arr)-1 || arr[mid+1] != target {
				return mid
			}
			left = mid + 1
		} else if arr[mid] < target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{10, 11, 14, 15}
	tg := 1
	fmt.Println(targetFirstLast(arr, tg))
}
