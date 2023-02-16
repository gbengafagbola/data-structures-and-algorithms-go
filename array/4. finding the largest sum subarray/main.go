package main
import "fmt"

// Naive solution of finding maximum sub array sum 
func MaxSubArraySum(data []int) int {
    max := data[0]
    sum := 0
    end := len(data) - 1

    for i := 0; i <= end; i++ {
        for j:=i; j <= end; j++ {
            sum += data[j]
            if sum > max {
                max = sum
            }
        } 
        sum = 0
    }
    return max
}

// Testing
func main() {
	arr := []int{1,-2,3,4,-4,6,-14}
	fmt.Println("MaxSubArraySum:", MaxSubArraySum(arr))
}

//  Time Complexity
// O(n^2)