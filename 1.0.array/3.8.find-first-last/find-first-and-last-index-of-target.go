package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1} 
	}

	left := 0
	right := len(nums) - 1
	found := false

	for left <= right {
		if nums[left] != target {
			left++
		}
		if nums[right] != target {
			right--
		}
		if left <= right && nums[left] == target && nums[right] == target {
			found = true
			break
		}
	}

	if found {
		return []int{left, right}
	}
	return []int{-1, -1} 
}

func searchRange2(nums []int, target int) []int {
	first := findFirst(nums, target)
	last := findLast(nums, target)
	return []int{first, last}
}

func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1 // Fix: Correct update when nums[mid] > target
		}
	}
	return -1
}

func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1 // Fix: Correct update when nums[mid] > target
		}
	}
	return -1
}


func main() {
	nums := []int{10, 11, 14, 15}
	tg := 1
	fmt.Println(searchRange(nums, tg))
}
