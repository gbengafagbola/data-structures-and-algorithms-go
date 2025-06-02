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