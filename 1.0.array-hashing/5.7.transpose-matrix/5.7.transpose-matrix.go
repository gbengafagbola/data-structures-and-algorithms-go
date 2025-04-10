package main

func TransposeMatrix(matrix [][]int) [][]int {
    if len(matrix) == 0 {
        return matrix
    }

    matrixMaxArray := len(matrix[0])
    result := make([][]int, matrixMaxArray)

    for _, row := range matrix {
        for i, v := range row {
            result[i] = append(result[i], v)
        }
    }

    return result
}