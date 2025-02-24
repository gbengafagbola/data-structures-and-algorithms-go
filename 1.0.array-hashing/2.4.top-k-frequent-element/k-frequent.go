package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for _, num := range nums {
		count[num]++
	}

	for num, cnt := range count {
		freq[cnt] = append(freq[cnt], num)
	}

	res := []int{}
	for i := len(freq) - 1; i > 0; i-- {
		for _, num := range freq[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}

func topKFrequentTwo(nums []int, k int) []int {
	frequencyMap := make(map[int]int)
	result := []int{}
	// Create a slice of key-value pairs
	type kv struct {
		Key   int
		Value int
	}

	for _, num := range nums {
		frequencyMap[num]++
	}

	// creates an unordered map
	var sortedSlice []kv
	for key, value := range frequencyMap {
		sortedSlice = append(sortedSlice, kv{Key: key, Value: value})
	}

	// Sort the slice by values in descending order
	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[i].Value > sortedSlice[j].Value // Change to `<` for ascending
	})

	for _, kv := range sortedSlice {
		result = append(result, kv.Key)
	}

	return result[0:k]
}

func main() {
	strs := []int{1, 1, 1, 2, 2, 3}
	k := 2
	result := topKFrequent(strs, k)
	fmt.Println(result)
}
