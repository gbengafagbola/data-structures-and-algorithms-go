package main
import "fmt"

func Sort1toN(arr []int, size int) {
    i := 0;
    for i < size {
        if arr[i] != i+1 {
            arr[i], arr[arr[i]-1] = arr[arr[i]-1], arr[i]
        } 
        if arr[i] == i+1{
           i++;
        }
    }
}

func main() {
	array := []int{ 8, 5, 6, 1, 9, 3, 2, 7, 4, 10}
	size := len(array);
	Sort1toN(array, size)
	fmt.Println("array", array)
}

//  Time Complexity
// O(n)