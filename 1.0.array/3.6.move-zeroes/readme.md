# Move Zeroes to End

## Problem Statement
Given an array of integers, move all the zeroes to the end while maintaining the relative order of the non-zero elements.

---

## Algorithm Explanation

The function `moveZeroes(nums []int) []int` efficiently moves all zeroes to the end using a **two-pointer approach**:

1. Initialize `zeroIndex = 0` to track the position where the next non-zero element should be placed.
2. Iterate over the array:
   - If the current element is non-zero, move it to the `zeroIndex` position and increment `zeroIndex`.
3. After processing all non-zero elements, fill the remaining positions with `0`s.

---

## Code Implementation

```go
package main

import "fmt"

func moveZeroes(nums []int) []int {
	zeroIndex := 0
	n := len(nums)

	for nonZeroIndex := 0; nonZeroIndex < n; nonZeroIndex++ {
		if nums[nonZeroIndex] != 0 {
			nums[zeroIndex] = nums[nonZeroIndex]
			zeroIndex++
		}
	}
	for i := zeroIndex; i < n; i++ {
		nums[i] = 0
	}

	return nums
}

func main() {
	nums := []int{2, 1, 0, 4, 0, 0, 6}
	result := moveZeroes(nums)
	fmt.Println(result)
}
```

---

## Example Walkthrough
### Input:
```go
nums = [2, 1, 0, 4, 0, 0, 6]
```

### Step-by-step Execution:
| Step | Array State | zeroIndex |
|------|------------|------------|
| Start | `[2, 1, 0, 4, 0, 0, 6]` | 0 |
| `2` (non-zero) → Swap | `[2, 1, 0, 4, 0, 0, 6]` | 1 |
| `1` (non-zero) → Swap | `[2, 1, 0, 4, 0, 0, 6]` | 2 |
| `0` (skip) | `[2, 1, 0, 4, 0, 0, 6]` | 2 |
| `4` (non-zero) → Swap | `[2, 1, 4, 4, 0, 0, 6]` | 3 |
| `0` (skip) | `[2, 1, 4, 4, 0, 0, 6]` | 3 |
| `0` (skip) | `[2, 1, 4, 4, 0, 0, 6]` | 3 |
| `6` (non-zero) → Swap | `[2, 1, 4, 6, 0, 0, 6]` | 4 |
| Fill remaining with `0`s | `[2, 1, 4, 6, 0, 0, 0]` | - |

### **Final Output:**
```go
[2, 1, 4, 6, 0, 0, 0]
```

---

## Complexity Analysis
- **Time Complexity:** `O(n)` - Iterates through the array twice (once for movement, once for filling zeroes).
- **Space Complexity:** `O(1)` - Modifies the array in place.

---

## Edge Cases Considered
| Case | Example Input | Expected Output |
|------|--------------|----------------|
| All zeroes | `[0, 0, 0, 0]` | `[0, 0, 0, 0]` |
| No zeroes | `[1, 2, 3, 4]` | `[1, 2, 3, 4]` |
| Mixed case | `[0, 1, 0, 3, 12]` | `[1, 3, 12, 0, 0]` |

---

## Summary
This implementation efficiently moves zeroes to the end while preserving the order of non-zero elements using an **in-place two-pointer technique**.

