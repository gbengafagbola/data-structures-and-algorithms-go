package main

import "fmt"

// Main function to get river sizes
func RiverSizes(matrix [][]int) []int {
	visited := make([][]bool, len(matrix))
	for i := range matrix {
		visited[i] = make([]bool, len(matrix[i]))
	}

	sizes := []int{}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if visited[i][j] {
				continue
			}
			traverseNode(i, j, matrix, visited, &sizes)
		}
	}
	return sizes
}

func traverseNode(i, j int, matrix [][]int, visited [][]bool, sizes *[]int) {
	currentRiverSize := 0
	nodesToExplore := [][]int{{i, j}}

	for len(nodesToExplore) > 0 {
		currentNode := nodesToExplore[len(nodesToExplore)-1]
		nodesToExplore = nodesToExplore[:len(nodesToExplore)-1]

		i, j := currentNode[0], currentNode[1]
		if visited[i][j] {
			continue
		}
		visited[i][j] = true

		if matrix[i][j] == 0 {
			continue
		}

		currentRiverSize++
		unvisitedNeighbors := getUnvisitedNeighbors(i, j, matrix, visited)
		for _, neighbor := range unvisitedNeighbors {
			nodesToExplore = append(nodesToExplore, neighbor)
		}
	}

	if currentRiverSize > 0 {
		*sizes = append(*sizes, currentRiverSize)
	}
}

func getUnvisitedNeighbors(i, j int, matrix [][]int, visited [][]bool) [][]int {
	unvisitedNeighbors := [][]int{}

	if i > 0 && !visited[i-1][j] {
		unvisitedNeighbors = append(unvisitedNeighbors, []int{i - 1, j})
	}
	if i < len(matrix)-1 && !visited[i+1][j] {
		unvisitedNeighbors = append(unvisitedNeighbors, []int{i + 1, j})
	}
	if j > 0 && !visited[i][j-1] {
		unvisitedNeighbors = append(unvisitedNeighbors, []int{i, j - 1})
	}
	if j < len(matrix[0])-1 && !visited[i][j+1] {
		unvisitedNeighbors = append(unvisitedNeighbors, []int{i, j + 1})
	}

	return unvisitedNeighbors
}

func main() {
	matrix := [][]int{
		{1, 0, 0, 1, 0},
		{1, 0, 1, 0, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 1, 0},
	}

	fmt.Println(RiverSizes(matrix)) // Example output: [2 1 5 2 2]
}
