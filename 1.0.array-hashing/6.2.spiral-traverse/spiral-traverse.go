package main  

import (
	"fmt"
)

func spiralTraverse(array [][]int) []int {
	result := []int{}
	if len(array) == 0 || len(array[0]) == 0 {
		return result
	}

	startRow, endRow := 0, len(array)-1
	startCol, endCol := 0, len(array[0])-1

	for startRow <= endRow && startCol <= endCol {
		// Traverse right
		for col := startCol; col <= endCol; col++ {
			result = append(result, array[startRow][col])
		}

		// Traverse down
		for row := startRow + 1; row <= endRow; row++ {
			result = append(result, array[row][endCol])
		}

		// Traverse left
		if startRow < endRow {
			for col := endCol - 1; col >= startCol; col-- {
				result = append(result, array[endRow][col])
			}
		}

		// Traverse up
		if startCol < endCol {
			for row := endRow - 1; row > startRow; row-- {
				result = append(result, array[row][startCol])
			}
		}

		startRow++
		endRow--
		startCol++
		endCol--
	}

	return result
}

func main() {
	array := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralTraverse(array)) // Output: [1 2 3 4 8 12 11 10 9 5 6 7]
}
