package main

import (
	"fmt"
	"math"
)

func lengthOfLongestSubString(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	left, right := 0, 0
	maxLen := 0

	chars := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if val, ok := chars[s[i]]; ok {
			left = int(math.Max(float64(left), float64(val+1)))//or implement a < b func
		}
		right++
		chars[s[i]] = i
		maxLen = int(math.Max(float64(right-left), float64(maxLen)))
	}

	return maxLen
}

func main() {
	s := "abcabcbb"

	result := lengthOfLongestSubString(s)
	fmt.Println(result)
}
