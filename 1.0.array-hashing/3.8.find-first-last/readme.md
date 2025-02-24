# Search Range Implementation in Go

## Overview
This implementation finds the **first and last position** of a target element in a sorted array using **binary search**. The algorithm efficiently locates the target in **O(log n)** time complexity.

## Function Breakdown

### `searchRange(nums []int, target int) []int`
- Calls `findFirst` to locate the first occurrence of `target`.
- Calls `findLast` to locate the last occurrence of `target`.
- Returns an array `[first, last]`. If `target` is not found, returns `[-1, -1]`.

```go
func searchRange(nums []int, target int) []int {
	first := findFirst(nums, target)
	last := findLast(nums, target)
	return []int{first, last}
}
```

---

### `findFirst(nums []int, target int) int`
- Uses **binary search** to find the first occurrence of `target`.
- If `nums[mid] == target`, checks if `mid` is the first occurrence.
- If `nums[mid] < target`, moves `left` pointer to `mid + 1`.
- Otherwise, moves `right` pointer to `mid - 1`.

```go
func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
```

#### **Illustration**
**Example:** `nums = [2, 4, 4, 4, 6, 8]`, `target = 4`

| Iteration | left | mid | right | nums[mid] | Action |
|-----------|------|-----|-------|-----------|---------|
| 1         | 0    | 2   | 5     | 4         | Move `right` to `mid - 1` |
| 2         | 0    | 1   | 1     | 4         | Found first occurrence at index `1` |

---

### `findLast(nums []int, target int) int`
- Uses **binary search** to find the last occurrence of `target`.
- If `nums[mid] == target`, checks if `mid` is the last occurrence.
- If `nums[mid] > target`, moves `right` pointer to `mid - 1`.
- Otherwise, moves `left` pointer to `mid + 1`.

```go
func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
```

#### **Illustration**
**Example:** `nums = [2, 4, 4, 4, 6, 8]`, `target = 4`

| Iteration | left | mid | right | nums[mid] | Action |
|-----------|------|-----|-------|-----------|---------|
| 1         | 0    | 2   | 5     | 4         | Move `left` to `mid + 1` |
| 2         | 3    | 3   | 5     | 4         | Found last occurrence at index `3` |

---

## **Time Complexity Analysis**
| Function      | Complexity |
|--------------|------------|
| `findFirst`  | O(log n)  |
| `findLast`   | O(log n)  |
| `searchRange`| O(log n)  |

### **Why O(log n)?**
- Binary search halves the search space at each step.
- Worst case: It takes `log₂(n)` iterations to locate `target`.

---

## **Example Test Cases**
```go
fmt.Println(searchRange([]int{}, 0)) // Expected: [-1, -1] (Empty array)
fmt.Println(searchRange([]int{1, 2, 3, 4, 5}, 6)) // Expected: [-1, -1] (Target not found)
fmt.Println(searchRange([]int{1, 1, 1, 1, 1}, 1)) // Expected: [0, 4] (All elements match target)
fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8)) // Expected: [3, 4] (Target found in middle)
fmt.Println(searchRange([]int{2, 2, 2, 3, 4, 5, 6, 7, 8, 8, 8, 8, 9}, 8)) // Expected: [8, 11] (Multiple occurrences)
fmt.Println(searchRange([]int{1, 3, 3, 5, 5, 5, 6, 8, 9}, 3)) // Expected: [1, 2] (Target at start)
fmt.Println(searchRange([]int{3, 3, 3, 3, 3, 3}, 3)) // Expected: [0, 5] (All are target)
```

---

## **Conclusion**
- **Binary search** ensures efficiency.
- **Edge cases** (empty array, no target, all elements matching) are covered.
- **Illustrations** help visualize execution flow.

[source](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array)