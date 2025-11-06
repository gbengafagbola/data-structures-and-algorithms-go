package main

import (
	"fmt"
	"math"
	"sort"
)

func SweetAndSavory(dishes []int, target int) []int {
	// Separate sweet and savory dishes
	sweetDishes := []int{}
	savoryDishes := []int{}

	for _, dish := range dishes {
		if dish < 0 {
			sweetDishes = append(sweetDishes, dish)
		} else if dish > 0 {
			savoryDishes = append(savoryDishes, dish)
		}
	}

	// Sort sweet dishes by absolute value and savory dishes normally
	sort.Slice(sweetDishes, func(i, j int) bool {
		return int(math.Abs(float64(sweetDishes[i]))) < int(math.Abs(float64(sweetDishes[j])))
	})
	sort.Ints(savoryDishes)

	bestPair := []int{0, 0}
	bestDifference := math.Inf(1)

	sweetIndex, savoryIndex := 0, 0

	for sweetIndex < len(sweetDishes) && savoryIndex < len(savoryDishes) {
		currentSum := sweetDishes[sweetIndex] + savoryDishes[savoryIndex]

		if currentSum <= target {
			currentDifference := float64(target - currentSum)
			if currentDifference < bestDifference {
				bestDifference = currentDifference
				bestPair = []int{sweetDishes[sweetIndex], savoryDishes[savoryIndex]}
			}
			savoryIndex++
		} else {
			sweetIndex++
		}
	}

	return bestPair
}

func main() {
	fmt.Println(SweetAndSavory([]int{-3, -5, 1, 7}, 8))      // [-3, 7]
	fmt.Println(SweetAndSavory([]int{-10, -5, 0, 5, 9}, 10)) // [-5, 9]
	fmt.Println(SweetAndSavory([]int{-5, -2, -1}, 3))        // [0, 0]
	fmt.Println(SweetAndSavory([]int{-8, -6, 6, 12}, 6))     // [-6, 12]
	fmt.Println(SweetAndSavory([]int{-4, -2, 2, 3}, 1))      // [-2, 3]
}
