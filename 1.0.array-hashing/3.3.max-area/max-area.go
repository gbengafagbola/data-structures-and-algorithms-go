package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	max_area := 0
	left := 0
	right := len(height) - 1

	for left < right {
		h := float64(min(height[left], height[right]))
		width := float64(right - left)
		max_area = int(math.Max(float64(max_area), h*width))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max_area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := maxArea(height)
	fmt.Println(result) // Expected output: 49
}
