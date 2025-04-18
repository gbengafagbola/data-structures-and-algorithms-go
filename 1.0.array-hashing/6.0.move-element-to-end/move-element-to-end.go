package main

func MoveElementToEnd(array []int, toMove int) []int {
	result := []int{}
	toMoveList := []int{}

	for _, val := range array {
		if val == toMove {
			toMoveList = append(toMoveList, val)
		} else {
			result = append(result, val)
		}
	}

	return append(result, toMoveList...)
}
