package main

func hasSingleCycle(array []int) bool {
	nVisited := 0 
	currentIdx := 0
	for nVisited < len(array) {
		if nVisited > 0 && currentIdx == 0 {
			return false
		}
		nVisited++
		currentIdx = (currentIdx + array[currentIdx]) % len(array)
		if currentIdx < 0 {
			currentIdx += len(array)
		}
	}
	return currentIdx == 0 
}