package main

import "fmt"

func containsDuplicate(nums []int) bool {
	counts := make(map[int]struct{})
	for _, value := range nums {
		if _, exists := counts[value]; exists {
			return true
		}
		counts[value] = struct{}{} // Uses zero-space struct instead of bool
	}
	return false
}

func containsDuplicateTwo(nums []int) bool {
	counts := make(map[int]bool, 0)
	for _, value := range nums {
		if _, ok := counts[value]; ok {
			return true
		}
		counts[value] = true
	}
	return false
}


func containsDuplicateThree(data []int) bool {
	counts := make(map[int]int)	
	for _, value := range data {
		counts[value]++
		if counts[value] > 1 {
			return true
		}
	}
	return false
}     


func main() {
	nums := []int{1, 2, 3, 4, 5, 9, 7, 9}
	fmt.Println("the array returns:", containsDuplicate(nums))
}
