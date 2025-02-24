package main
import "fmt"

func indexArray(arr []int, size int) {
    i := 0;
    for i < size {
        if arr[i] >= 0 && i != arr[i] {
            arr[i], arr[arr[i]] = arr[arr[i]], arr[i]
        } 
        if arr[i] == i || arr[i] < 0{
           i++
        }
    }
}

// run / testing
func main() {
	array := []int{ 8, -1, 6, 1, 9, 3, 2, 7, 4, -1 }
	size := len(array);
	indexArray(array, size)
	fmt.Println("array", array)
}

//  Time Complexity
// O(n)
