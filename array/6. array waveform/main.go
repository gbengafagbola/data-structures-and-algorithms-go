package main
import "fmt"

func WaveArray(arr []int) {
	for i := 1; i < len(arr); i += 2 {
        if (i-1) >= 0 && arr[i] > arr[i-1] {
            arr[i], arr[i-1] = arr[i-1], arr[i]
        }
        if (i+1) < len(arr) && arr[i] > arr[i+1] {
            arr[i], arr[i+1] = arr[i+1], arr[i]
        }
	}
}

func main() {
	array := []int{8, 1, 2, 3, 4, 5, 6, 4, 2}
	WaveArray(array)
	fmt.Println("array", array)
}


//  Time Complexity
// O(n)