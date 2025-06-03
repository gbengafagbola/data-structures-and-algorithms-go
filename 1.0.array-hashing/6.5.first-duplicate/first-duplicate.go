package main

func FirstDuplicateValue(array []int) int {
	n := make(map[int]bool)

	for _, value := range array {
		if _, exists := n[value]; exists {
			return value // Return the first duplicate found
		}
		n[value] = true // Mark the value as seen
	}
	return -1 // Return -1 if no duplicate is found
}


//since we can mutate the array and all integers are in the range of 1 to n, we can use the array itself to track seen values
func FirstDuplicateValue2(array []int) int {
	for i := 0; i < len(array); i++ {
		index := abs(array[i]) - 1 // Get the index corresponding to the value
		if array[index] < 0 {
			return abs(array[i]) // If the value at that index is negative, we have seen this value before
		}
		array[index] = -array[index] // Mark the value as seen by negating it
	}
	return -1 // Return -1 if no duplicate is found
}


func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}