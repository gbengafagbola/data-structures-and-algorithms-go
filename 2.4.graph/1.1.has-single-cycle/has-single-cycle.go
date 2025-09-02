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




func HasSingleCycleHelperFunc(array []int) bool {
	nVisitedElement := 0 
	currentIdx := 0
	for nVisitedElement <  len(array) {
		if nVisitedElement > 0 && currentIdx == 0 {
			return false
		}
		nVisitedElement++
		currentIdx = getNextIdx(currentIdx, array)
	}
	return currentIdx == 0 
}

func getNextIdx(currentIdx int, array []int) int {
    
    jump := array[currentIdx]
    nextIdx := (currentIdx + jump) % len(array)
    
    if nextIdx >= 0 {
        return nextIdx
    } else {
      return  nextIdx + len(array)
    }
    
}