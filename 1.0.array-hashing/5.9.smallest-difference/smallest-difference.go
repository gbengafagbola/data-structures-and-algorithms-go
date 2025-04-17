package main

import (
    "math"
    "sort"
)

func SmallestDifference(array1, array2 []int) []int {
    sort.Ints(array1)
    sort.Ints(array2)
    
    result := []int{0, 0}
    smallest := math.MaxInt64
    idxAone := 0
    idxAtwo := 0

    for idxAone < len(array1) && idxAtwo < len(array2) {
        pair1 := array1[idxAone]
        pair2 := array2[idxAtwo]
        
        // Calculate absolute difference
        currentDiff := int(math.Abs(float64(pair1 - pair2)))
        
        // Update result if current difference is smaller
        if currentDiff < smallest {
            smallest = currentDiff
            result = []int{pair1, pair2}
        }
        
        // Move the pointer for the smaller number
		
        if pair1 < pair2 {
            idxAone++
        } else if pair2 < pair1 {
            idxAtwo++
        } else {
            // If equal, move both pointers (optional optimization)
            idxAone++
            idxAtwo++
        }
    }

    return result
}

func SmallestDifference2(array1, array2 []int) []int {
	smallest := math.MaxInt64
	result := []int{}

	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			diff := int(math.Abs(float64(array1[i] - array2[j])))
			if diff < smallest {
				smallest = diff
				result = []int{array1[i], array2[j]}
			}
		}
	}

	return result
}
