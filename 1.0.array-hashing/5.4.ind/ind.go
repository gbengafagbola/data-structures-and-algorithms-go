package main

import "fmt"

func numberOfItems(s string, startIndices, endIndices []int32) []int32 {
	n := len(s)
	prefixSum := make([]int, n)  // Stores the cumulative count of `*`
	pipeIndices := make([]int, 0) // Stores indices of `|` pipes

	// Build prefix sums and store positions of `|`
	count := 0
	for i := 0; i < n; i++ {
		if s[i] == '*' {
			count++
		}
		prefixSum[i] = count

		if s[i] == '|' {
			pipeIndices = append(pipeIndices, i)
		}
	}

	result := make([]int32, len(startIndices))

	// Process each query
	for i := 0; i < len(startIndices); i++ {
		start := int(startIndices[i]) - 1
		end := int(endIndices[i]) - 1

		// Find the first `|` on or after `start`
		left := -1
		for _, idx := range pipeIndices {
			if idx >= start {
				left = idx
				break
			}
		}

		// Find the last `|` on or before `end`
		right := -1
		for j := len(pipeIndices) - 1; j >= 0; j-- {
			if pipeIndices[j] <= end {
				right = pipeIndices[j]
				break
			}
		}

		// Count stars (`*`) between `|` boundaries
		if left != -1 && right != -1 && left < right {
			result[i] = int32(prefixSum[right] - prefixSum[left])
		} else {
			result[i] = 0
		}
	}

	return result
}

func numberOfItems2(s string, startIndices, endIndicies []int32) []int32 {
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

	fmt.Println(numberOfItems(s, startIndices, endIndices)) // Output: [2, 3]
}
