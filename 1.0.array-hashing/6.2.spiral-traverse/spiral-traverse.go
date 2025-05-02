package main  

import (
	"fmt"
)

func spiralTraverse(matrix [][]int) []int {
    result := []int{}
    if len(matrix) == 0 {
        return result
    }

    top, bottom := 0, len(matrix)-1
    left, right := 0, len(matrix[0])-1

    for top <= bottom && left <= right {
        for col := left; col <= right; col++ {
            result = append(result, matrix[top][col])
        }
        top++

        for row := top; row <= bottom; row++ {
            result = append(result, matrix[row][right])
        }
        right--

        if top <= bottom {
            for col := right; col >= left; col-- {
                result = append(result, matrix[bottom][col])
            }
            bottom--
        }

        if left <= right {
            for row := bottom; row >= top; row-- {
                result = append(result, matrix[row][left])
            }
            left++
        }
    }

    return result
}


func spiralTraverseRecursive(matrix [][]int) []int {
    result := []int{}
    if len(matrix) == 0 {
        return result
    }
    spiralFill(matrix, 0, len(matrix)-1, 0, len(matrix[0])-1, &result)
    return result
}

func spiralFill(matrix [][]int, startRow, endRow, startCol, endCol int, result *[]int) {
    if startRow > endRow || startCol > endCol {
        return
    }

    for col := startCol; col <= endCol; col++ {
        *result = append(*result, matrix[startRow][col])
    }

    for row := startRow + 1; row <= endRow; row++ {
        *result = append(*result, matrix[row][endCol])
    }

    if startRow < endRow {
        for col := endCol - 1; col >= startCol; col-- {
            *result = append(*result, matrix[endRow][col])
        }
    }

    if startCol < endCol {
        for row := endRow - 1; row > startRow; row-- {
            *result = append(*result, matrix[row][startCol])
        }
    }

    spiralFill(matrix, startRow+1, endRow-1, startCol+1, endCol-1, result)
}

func main() {
	array := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralTraverse(array)) // Output: [1 2 3 4 8 12 11 10 9 5 6 7]
}
