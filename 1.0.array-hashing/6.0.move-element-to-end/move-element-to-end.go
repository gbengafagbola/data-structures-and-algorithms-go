package main

func MoveElementToEnd(array []int, toMove int) []int {
	i := 0
	j := len(array) -1

	for i < j {
		if array[j] != toMove && array[i] == toMove {
			array[i], array[j] = array[j], array[i]
			i++
			j--
		} else if array[j] == toMove && array[i] == toMove {
			j--
		} else {
			i++
		}
	}
	return array
}



func MoveElementToEnd2(array []int, toMove int) []int {
	left := 0
	right := len(array) - 1

	for left < right {
		for left < right && array[right] == toMove {
			right--
		}
		if array[left] == toMove {
			array[left], array[right] = array[right], array[left]
		}
		left++
	}

	return array
}

func MoveElementToEnd3(array []int, toMove int) []int {
	result := []int{}
	toMoveList := []int{}

	for _, val := range array {
		if val == toMove {
			toMoveList = append(toMoveList, val)
		} else {
			result = append(result, val)
		}
	}
	// spreading individual elements of toMoveList into result, because you can't just merge a slice into another slice so you have to offload one by one
	return append(result, toMoveList...)
}
