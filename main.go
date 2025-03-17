// package main

// import (
// 	"dsa-golang/generate"
// 	"log" // Importing log package
// )

// func main() {
// 	// Call the Generate function and handle errors
// 	if err := generate.Generate(); err != nil {
// 		log.Fatalf("Error: %v", err)  // Log the error and exit
// 	}
// }

package main
import "fmt"

func numberOfItems(s string, startIndices, endIndicies []int32) []int32 {
	subString1 := s[((startIndices[0]) -1) : (endIndicies[0])] // s[1:5]
	subString2 := s[((startIndices[1]) -1) : (endIndicies[1])] // s[1:6]

	count1 := 0
	count2 := 0

	a1 := []int{}
	a2 := []int{}

	for i := 0; i < len(subString1); i++ {
		if (subString1[i]-'a') == 27 {
			a1 = append(a1, i)
			if len(a1) == 2{
				count1 += ((a1[1] - a1[0]) -1)
				a1 = []int{a1[1]}
			}
		}
	}

	for i := 0; i < len(subString2); i++ {
		if(subString2[i]-'a') == 27 {
			a2 = append(a2, i)
			if len(a2) == 2{
				count2 += ((a2[1] - a2[0]) -1)
				a2 = []int{a2[1]}
			}
		}
	}

	answer := []int32{int32(count1), int32(count2)}
	return answer
}

func main() {
	startIndices := []int32{1, 1}
	endIndices := []int32{5, 6}

	s := "|**|*|*"

	fmt.Println(numberOfItems(s, startIndices, endIndices))
}