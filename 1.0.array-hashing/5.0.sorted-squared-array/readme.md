Let's break down the optimized **SortedSquaredArray** function visually and step by step.

---

### **Understanding the Problem**
We have a sorted array of integers (both negative and positive), and we want to return a new array where each number is squared and the result remains sorted.

#### **Example**
```go
input:  [-7, -3, 1, 4, 8]
output: [1, 9, 16, 49, 64]
```
Simply squaring all values gives us `[49, 9, 1, 16, 64]`, which isn't sorted. Sorting afterward would take **O(n log n)** time, which isn't optimal.

---

### **Optimized Approach: Two-Pointer Technique**
We use **two pointers**, one starting at the leftmost (`left` pointer) and the other at the rightmost (`right` pointer), because the **largest absolute values** are either at the beginning or the end.

#### **Step-by-step Explanation**
##### **Step 1: Initialize Pointers**
```go
left, right := 0, n-1 // left at the beginning, right at the end
```
For `array = [-7, -3, 1, 4, 8]`
```
  L                     R
[-7, -3,  1,  4,  8]
```
We also initialize `result = [_, _, _, _, _]` (empty result array of size `n`).

##### **Step 2: Compare Absolute Values**
We compare `abs(array[left])` and `abs(array[right])`:
- `| -7 | = 7`
- `|  8 | = 8`
Since `8` is greater, square it and place it at the **end** of `result`:
```go
result[4] = 8² = 64
```
Move `right` pointer leftward (`right--`).
```
  L                 R
[-7, -3,  1,  4,  8]   => result = [_, _, _, _, 64]
```

##### **Step 3: Repeat the Process**
Now compare again:
- `| -7 | = 7`
- `|  4 | = 4`
Since `7 > 4`, square `7` and insert:
```go
result[3] = 7² = 49
```
Move `left` pointer rightward (`left++`).
```
      L             R
[-7, -3,  1,  4,  8]   => result = [_, _, _, 49, 64]
```

##### **Step 4: Continue Processing**
- Compare `|-3| = 3` and `|4| = 4`
- `4 > 3`, so square `4`:
```go
result[2] = 4² = 16
```
Move `right` pointer leftward (`right--`).
```
      L         R
[-7, -3,  1,  4,  8]   => result = [_, _, 16, 49, 64]
```

- Compare `|-3| = 3` and `|1| = 1`
- `3 > 1`, so square `-3`:
```go
result[1] = (-3)² = 9
```
Move `left++`.
```
         L      R
[-7, -3,  1,  4,  8]   => result = [_, 9, 16, 49, 64]
```

- Square `1`:
```go
result[0] = 1² = 1
```
Final result:
```
[1, 9, 16, 49, 64]
```

---

### **Final Code Recap**
```go
func SortedSquaredArray(array []int) []int {
	n := len(array)
	result := make([]int, n)
	left, right := 0, n-1

	for i := n - 1; i >= 0; i-- {
		if abs(array[left]) > abs(array[right]) {
			result[i] = array[left] * array[left]
			left++
		} else {
			result[i] = array[right] * array[right]
			right--
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

---

### **Time & Space Complexity**
✅ **O(n) Time Complexity** - We process each element **once**.  
✅ **O(n) Space Complexity** - We use a new array `result` of the same size as input.

Would you like a graphical animation of the sorting process? 🚀