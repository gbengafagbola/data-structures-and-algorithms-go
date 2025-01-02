package main

import (
	"fmt"
	"sort"
)

// func topKFrequent(nums []int, k int) []int {
// 	frequencyMap := make(map[int]int)
// 	result := []int{}

// 	for _, num := range nums {
// 		frequencyMap[num]++
// 	}

// 	for _, value := range frequencyMap {
// 		if value <= k {
// 			result = append(result, value)
// 		}
// 	}

// 	sort.Slice(result, func(i, j int) bool {
//         return result[i] > result[j] // Change to `<` for ascending
//     })

// 	return result[:k]
// }

func topKFrequent(nums []int, k int) []int {
	frequencyMap := make(map[int]int)
	result := []int{}
	final_result := []int{}

	for _, num := range nums {
		frequencyMap[num]++
	}

	for _, value := range frequencyMap {
		result = append(result, value)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j] // Change to `<` for ascending
	})

	results := result[:k]

	for _, result := range results {
		final_result =  append(final_result, frequencyMap[result])
	}

	return final_result
}

func main() {
	strs := []int{1, 1, 1, 2, 2, 3}
	k := 2
	result := topKFrequent(strs, k)
	fmt.Println(result)
	// fmt.Println(fre)
}
