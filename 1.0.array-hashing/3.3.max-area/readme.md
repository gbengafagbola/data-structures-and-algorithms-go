# Container With Most Water

## Problem Statement
Given an array `height[]` where each element represents the height of a vertical line, find two lines that together with the x-axis form a container, such that the container holds the most water.

## Algorithm Explanation
We use the **two-pointer technique**:
1. **Initialize two pointers**: One at the start (`left = 0`) and one at the end (`right = len(height) - 1`).
2. **Compute the area** using the formula:
   
   \[ \text{Area} = \min(\text{height}[\text{left}], \text{height}[\text{right}]) \times (\text{right} - \text{left}) \]

3. **Move the pointer with the smaller height** to potentially find a larger area.
4. **Repeat until left and right pointers meet.**

---
## Code Implementation

```go
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
	fmt.Println("Maximum Water Area:", result)
}
```

---
## Visual Representation

Given an array:
```
Index:  0  1  2  3  4  5  6  7  8
Heights: 1  8  6  2  5  4  8  3  7
```

A possible container is:
```
   |                  |
   |        |        |
   |        |   |    |
   |    |   |   |    |
   |    |   |   | |  |
   |  | |   |   | |  |
   |__|_|___|___|_|__|
```
The two tallest lines are at index `1` and `8`, giving the maximum area:

\[ 8 \times (8 - 1) = 49 \]

---
## Complexity Analysis
- **Time Complexity**: \(O(n)\) - Since we traverse the list only once.
- **Space Complexity**: \(O(1)\) - Uses only a few extra variables.

---
## Example Usage

### Input:
```go
height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
```

### Output:
```shell
Maximum Water Area: 49
```

---
## Summary
This solution efficiently finds the largest container that can hold water using the two-pointer technique. The approach works in **linear time** and ensures optimal performance.

Happy Coding! 🚀

