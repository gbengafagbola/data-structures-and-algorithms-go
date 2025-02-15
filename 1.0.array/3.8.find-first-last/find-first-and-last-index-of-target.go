package main
import "fmt"

func target(arr []int, target int) []int {

	// if len(arr) <= 1 {

	// }

	left := 0 
	right := len(arr)-1

	for left <= right {

		if arr[left] != target {
			left++
		}

		if arr[right] != target {
			right--
		}

		if arr[left] == target && arr[right] == target{
			return []int{left, right}
		}

	}
	return []int{-1, -1}
}

func main() {
	arr := []int{10, 11, 14, 15}
	tg := 1
	fmt.Println(target(arr, tg))
}