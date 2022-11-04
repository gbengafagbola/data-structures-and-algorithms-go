package main
import ("fmt")

func SumArray(data []int) int { 
	total := 0
	  
	//Implement your solution here
	for i := 0; i < len(data); i++ {
	  total+= data[i]
	}
	
	return total
  }

//Testing code
func main() {
    data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Println("Sum of all the values in array:", SumArray(data)) 
}

//  Time Complexity
// O(n)