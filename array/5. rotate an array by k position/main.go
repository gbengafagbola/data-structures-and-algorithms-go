package main
import "fmt"

func RotateArray (data []int, k int){
	arr1 := data[:k]
	arr2 := data[k:]

	arr2 = append(arr2, arr1...)

	for i := 0; i < len(data); i++ {
		data[i] = arr2[i]
	}
}


func main() {
	data := []int{1, 2, 3, 4, 5, 6}
	RotateArray(data, 2)
	fmt.Println(data)
}