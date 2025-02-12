package main

import (
	"fmt"
	"sort"
)

// Function to calculate the minimum boats required
func rescueBoatsRequired(people []int, limit int) int {
	sort.Ints(people) // Sort people by weight

	smallest := 0
	biggest := len(people) - 1
	boats := 0

	for smallest <= biggest {
		// If the lightest and heaviest person fit in one boat
		if people[smallest]+people[biggest] <= limit {
			smallest++
			biggest--
		} else {
			// Otherwise, the heaviest person goes alone
			biggest--
		}
		boats++ // Increase boat count
	}

	return boats
}

func main() {
	people := []int{3, 2, 4, 5, 1}
	limit := 6

	result := rescueBoatsRequired(people, limit)
	fmt.Println(result) // Output: 3
}
