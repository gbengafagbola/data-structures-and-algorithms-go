package main
import ("fmt")

func SequentialSearch(data []int, value int) bool { 
    for i := 0; i < len(data); i++ {
        if data[i] == value {
            return true;
        } 
    }
    return false 
}

//Testing code
func main() {
  arr := []int{1, 2, 3, 4, 5, 6, 7, 9} 
  fmt.Println("SequentialSearch:", SequentialSearch(arr, 8))
  fmt.Println("SequentialSearch:", SequentialSearch(arr, 9))
}

//  Time Complexity
// O(n)