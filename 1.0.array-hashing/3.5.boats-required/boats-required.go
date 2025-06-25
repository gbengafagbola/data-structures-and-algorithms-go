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


func rescueBoatsRequired2(people []int, limit int) int {
	count := 0
	seen := make(map[int]bool)

	for _, v := range people {
		if v == limit {
			count++ // single person takes a boat
		} else if seen[limit - v] {
			count++ // found a pair
			delete(seen, limit - v)
		} else {
			seen[v] = true // mark this weight for future pairing
		}
	}

	// Remaining unmatched people each need a boat
	for range seen {
		count++
	}

	return count
}

func main() {
	people := []int{3, 2, 4, 5, 1}
	limit := 6

	result := rescueBoatsRequired(people, limit)
	fmt.Println(result) // Output: 3
}
