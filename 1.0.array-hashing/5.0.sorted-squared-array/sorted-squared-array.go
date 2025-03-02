package main
import (
	"fmt"
	"sort"
)

func SortedSquaredArray(array []int) []int {
	n := len(array)
	result := make([]int, n)
	left, right := 0, n-1

	for i := n - 1; i >= 0; i-- {
		if abs(array[left]) > abs(array[right]) {
			result[i] = array[left] * array[left]
			left++
		} else {
			result[i] = array[right] * array[right]
			right--
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SortedSquaredArray2(array []int) []int {
    result := []int{}

    for _, num := range array {
       i := num * num
        result = append(result, i)
    }

	 sort.Ints(result)
	
    return result
}


func main(){
	array:= []int{-3, -2, -1, 0, 1, 2, 3}
	fmt.Println(SortedSquaredArray(array))
}