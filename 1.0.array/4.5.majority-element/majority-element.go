package main

import "fmt"

func majorityElement(nums []int) int {
	check := len(nums) / 2
	tab := make(map[int]int)

	for _, num := range nums {
		tab[num]++ 

		if tab[num] > check { 
			return num
		}
	}

	return 0 
}

func main () {
	x := []int{1,3,2,1,1}
	fmt.Println(majorityElement(x))
}
