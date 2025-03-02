package main
import (
	"fmt"
	"sort"
)

func SortedSquaredArray(array []int) []int {
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