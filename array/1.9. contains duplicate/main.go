package main

import "fmt"

func containsDuplicate(data []int) bool {
	counts := make(map[int]int)
	for _, i := range data {
		counts[i]++
		if counts[i] > 1 {
			return true
		}
	}
	return false
}  

func containsDuplicate2(arr []int) bool {
	duplicateMap := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		// if a key dosen't exist in a map we get the value type's zero
		if duplicateMap[arr[i]] == 0 {
			duplicateMap[arr[i]] = 1
		} else {
			return false
		}
	}
	return true
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 9, 7, 9}
	fmt.Println("the array returns:", containsDuplicate(nums))
}
