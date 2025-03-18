package main

import (
	"fmt"
	"sort"
)

// [1 1 2 3 5 7 22]
func NonConstructibleChange(coins []int) int {
	sort.Ints(coins)
	change := 0

	for _, coin := range coins {
		if coin > change+1 {
			return change + 1
		}
		change += coin
	}

	return change + 1
}

func main() {
	coins := []int{1, 2, 5}
	fmt.Println(NonConstructibleChange(coins)) // Output: 4
}