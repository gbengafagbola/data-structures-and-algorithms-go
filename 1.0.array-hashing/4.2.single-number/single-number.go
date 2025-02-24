package main

import "fmt"

func singleNumber(nums []int) int {
	answer := 0
	for _, num := range nums {
		answer ^= num
	}
	return answer
}

func singleNumber2(nums []int) int {
	tab := make(map[int]int)

	for _, num := range nums {
		tab[num]++
	}

	for k, v := range tab {
		if v == 1 {
			return k
		}
	}

	return -1
}

func main() {
	nums := []int{4, 6, 2, 3, 4, 2, 3}
	fmt.Println(singleNumber(nums))
}
