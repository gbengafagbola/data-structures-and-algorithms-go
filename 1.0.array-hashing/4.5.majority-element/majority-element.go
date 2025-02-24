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



func majorityElement2(nums []int) int {
	count := 1
	majority := nums[0]
	n:= len(nums)
	for i := 1; i < n; i++ {
		if majority == nums[i] {
			count++
		} else {
			count--
			if count == 0 {
				majority = nums[i]
				count = 1
			}
		}
	}
	return majority
}

func main () {
	nums := []int{1,3,2,1,1}
	fmt.Println(majorityElement(nums))
}
