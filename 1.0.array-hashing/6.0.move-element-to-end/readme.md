Here's a very detailed `README.md` that explains both versions of your `MoveElementToEnd` function — the in-place version and the order-preserving version.

---

# 🧱 Move Element to End in Go

This Go package provides two approaches to solving a classic array manipulation problem: **"Move all instances of a given element to the end of an array."**

It’s a common task during coding interviews and is helpful for understanding array traversal, pointer manipulation, and memory optimization.

---

## 🧩 Problem Statement

Given an array of integers and a target integer `toMove`, move all instances of `toMove` to the end of the array. The order of the other elements does **not** have to be preserved (unless specifically required).

---

## 🧪 Example

```go
Input:
  array = [1, 2, 2, 2, 2, 1, 4, 3, 2]
  toMove = 2

Output:
  [1, 3, 4, 1, 2, 2, 2, 2, 2] // The 2s are at the end. Order of others may vary.
```

---

## ✅ Approaches

### 1. In-Place Swapping Using Two Pointers (Efficient)

```go
func MoveElementToEnd(array []int, toMove int) []int {
	left := 0
	right := len(array) - 1

	for left < right {
		for left < right && array[right] == toMove {
			right--
		}
		if array[left] == toMove {
			array[left], array[right] = array[right], array[left]
		}
		left++
	}
	return array
}
```

### 🧠 How It Works

- Use two pointers (`left` and `right`) from both ends of the array.
- If the `right` pointer is already pointing at `toMove`, just move it left.
- If the `left` pointer is at a `toMove` value, swap it with the `right` pointer's value.
- Continue this until the pointers meet.
- **Time Complexity**: O(n)
- **Space Complexity**: O(1) — done in-place

### 🧪 Example

```go
Input:  [2, 1, 2, 3, 2, 4]
Output: [4, 1, 3, 2, 2, 2]  // 2s moved to the end (order not guaranteed)
```

---

### 2. Order-Preserving With Extra Space (Clean)

```go
func MoveElementToEndPreservingOrder(array []int, toMove int) []int {
	result := []int{}
	toMoveList := []int{}

	for _, val := range array {
		if val == toMove {
			toMoveList = append(toMoveList, val)
		} else {
			result = append(result, val)
		}
	}

	return append(result, toMoveList...)
}
```

### 🧠 How It Works

- Create two new slices:
  - `result` for non-`toMove` values
  - `toMoveList` for `toMove` values
- Traverse the original array once.
- At the end, **concatenate** the slices using:
  
  ```go
  return append(result, toMoveList...)
  ```

- The `...` (variadic expansion operator) **unpacks** the slice `toMoveList` into individual elements, so it can be appended properly.

### 🧪 Example

```go
Input:  [2, 1, 2, 3, 2, 4]
Output: [1, 3, 4, 2, 2, 2]  // Preserves order of non-2 elements
```

### 🧠 Why the `...` Operator?

In Go, when you want to append one slice to another, you must unpack the second slice using `...`. Otherwise, Go will think you're trying to append a slice **as a single element**, which is not allowed.

```go
// Correct
append(result, toMoveList...)

// Incorrect
append(result, toMoveList) // compilation error
```

---

## 🚀 Example `main.go`

```go
package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 2, 2, 2, 1, 4, 3, 2}
	toMove := 2

	fmt.Println("In-Place:", MoveElementToEnd(array, toMove))

	array2 := []int{1, 2, 2, 2, 2, 1, 4, 3, 2}
	fmt.Println("Order-Preserving:", MoveElementToEndPreservingOrder(array2, toMove))
}
```

---

## 🔬 When to Use What?

| Use Case                              | Function Name                     | Memory Efficient | Preserves Order |
|--------------------------------------|----------------------------------|------------------|-----------------|
| In-place manipulation, best for large arrays | `MoveElementToEnd`               | ✅ Yes           | ❌ No            |
| You need to preserve order              | `MoveElementToEndPreservingOrder` | ❌ No (uses extra slices) | ✅ Yes         |

