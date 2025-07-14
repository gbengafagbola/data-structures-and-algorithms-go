# Valid Mountain Array

## Problem Statement
A valid mountain array is defined as an array that:
1. Has at least three elements.
2. Strictly increases to a single peak.
3. Strictly decreases after the peak.

The function `validMountainArray` checks if a given integer array follows the mountain shape.

---

## Algorithm Explanation

The function follows these steps:

1. Start from the second element and move forward as long as the sequence is increasing.
2. If the peak is the first or last element, return `false`.
3. Continue moving forward, checking if the sequence is strictly decreasing.
4. If we reach the end of the array, return `true`; otherwise, return `false`.

---

## Code Implementation
```go
package main

import "fmt"

func validMountainArray(arr []int) bool {
	index := 1

	// Walk up
	for index < len(arr) && arr[index] > arr[index-1] {
		index++
	}

	// breaking out from the for loop above indicate we have identified our peak and hence the Peak can't be the first or last element in the array thereby we return false
	if index == 1 || index == len(arr) {
		return false
	}

	// Walk down
	for index < len(arr) && arr[index] <= arr[index-1] {
		index++  
	}

	return index == len(arr)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	result := validMountainArray(arr)
	fmt.Print(result)
}
```

---

## Visual Flow

### Example Input
`arr = [1, 2, 3, 4, 5, 4, 3, 2, 1]`

### Step-by-step Execution
```
Index:   0  1  2  3  4  5  6  7  8
Array:   1  2  3  4  5  4  3  2  1
            ↗  ↗  ↗  ↗  ↑  ↘  ↘  ↘  ↘
```
- **Increasing Phase:** The loop moves up to `5` (index `4`).
- **Peak Check:** Ensures `5` is not at index `0` or `8`.
- **Decreasing Phase:** The loop moves down to `1` (index `8`).
- **Final Check:** Since the traversal reaches the end, return `true`.

### Edge Cases
#### **Not a Mountain Array**
| Case | Reason |
|------|--------|
| `[3, 5, 5]` | No strict increase or decrease |
| `[1, 2, 3]` | No downward slope |
| `[3, 2, 1]` | No upward slope |

---

## Complexity Analysis
- **Time Complexity:** `O(n)`, since we traverse the array at most twice.
- **Space Complexity:** `O(1)`, as we use only one variable (`index`).

---

## Summary
The function efficiently determines whether an array represents a mountain structure using a two-pass approach. It ensures a strict increase to a peak followed by a strict decrease, returning `true` only if both conditions are met.


[source](https://leetcode.com/problems/valid-mountain-array)
