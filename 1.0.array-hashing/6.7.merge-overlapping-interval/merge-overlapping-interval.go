package main

import (
	"fmt"
	"sort"
)

func MergeOverlappingIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// Step 1: Sort intervals by the start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}
	current := intervals[0]

	// Step 2: Merge overlapping intervals
	for i := 1; i < len(intervals); i++ {
		if current[1] >= intervals[i][0] {
			// Overlap, so merge
			current[1] = max(current[1], intervals[i][1])
		} else {
			// No overlap, add the previous one and update current
			result = append(result, current)
			current = intervals[i]
		}
	}
	result = append(result, current) // Add the last interval

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Example usage
func main() {
	input := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merged := MergeOverlappingIntervals(input)
	fmt.Println(merged) // Output: [[1 6] [8 10] [15 18]]
}
