package main

import "fmt"

//best performance
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

// v.bad perfomance
func containsDuplicate2(data []int) bool {
	duplicateMap := make(map[int]int)

	for i := 0; i < len(data); i++ {
		if duplicateMap[data[i]] == 0 {
			duplicateMap[data[i]] = 1
		} else {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 9, 7, 9}
	fmt.Println("the array returns:", containsDuplicate(nums))
}
