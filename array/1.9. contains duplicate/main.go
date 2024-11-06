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

func main() {
	nums := []int{1, 2, 3, 4, 5, 9, 7, 9}
	fmt.Println("the array returns:", containsDuplicate(nums))
}
