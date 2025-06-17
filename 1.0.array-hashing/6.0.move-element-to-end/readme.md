# 🧱 Move Element to End in Go

This Go package provides **three** different approaches to solving a classic array manipulation problem:  
**"Move all instances of a given element to the end of an array."**

This problem is frequently asked in technical interviews and is great for understanding pointer manipulation, time-space trade-offs, and in-place algorithms.

---

## 🧩 Problem Statement

> Given an array of integers and a target integer `toMove`, move all instances of `toMove` to the end of the array.

---

## 🧪 Example

```go
Input:
  array = [1, 2, 2, 2, 2, 1, 4, 3, 2]
  toMove = 2

Output:
  [1, 3, 4, 1, 2, 2, 2, 2, 2]
```

---

## ✅ Implementations

### 1. 🔁 In-Place with Two Pointers (Cleanest)

```go
func MoveElementToEnd(array []int, toMove int) []int {
	i := 0
	j := len(array) - 1

	for i < j {
		if array[j] != toMove && array[i] == toMove {
			array[i], array[j] = array[j], array[i]
			i++
			j--
		} else if array[j] == toMove && array[i] == toMove {
			j--
		} else {
			i++
		}
	}
	return array
}
```

#### 🔍 How It Works

- Starts with two pointers (`i` from the start, `j` from the end).
- If `i` points to `toMove` and `j` doesn’t, swap them.
- If both point to `toMove`, only move `j` left.
- Otherwise, increment `i`.
- **In-place**, efficient, and minimal branching.

#### ⏱️ Time & Space

- Time Complexity: **O(n)**
- Space Complexity: **O(1)**

---

### 2. ♻️ In-Place with Right Skip Optimization

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

#### 🔍 Difference from version 1:

- Optimizes by skipping consecutive `toMove`s from the end using an inner loop.
- Swaps immediately when it finds `toMove` on the left.

#### ⏱️ Time & Space

- Time Complexity: **O(n)**
- Space Complexity: **O(1)**

---

### 3. 🧼 Order-Preserving with Extra Space

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

#### 🧠 Why `...` (spread operator)?

The `...` tells Go to unpack the slice (`toMoveList`) and pass its elements individually into the `append` function. Without it, the compiler would raise an error.

```go
// ✅ Correct
append(result, toMoveList...)

// ❌ Error
append(result, toMoveList)
```

#### ⏱️ Time & Space

- Time Complexity: **O(n)**
- Space Complexity: **O(n)** (uses extra memory for two slices)

---

## 🔬 When Should You Use Each?

| Use Case                                | Function Name                       | Memory Efficient | Preserves Order |
|----------------------------------------|------------------------------------|------------------|-----------------|
| Best overall performance and clarity   | `MoveElementToEnd` (newest version) | ✅ Yes           | ❌ No            |
| Extra optimization on right skips      | `MoveElementToEnd` (optimized loop) | ✅ Yes           | ❌ No            |
| You need to preserve order             | `MoveElementToEndPreservingOrder`  | ❌ No (O(n))     | ✅ Yes           |

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

	fmt.Println("In-Place (Latest):", MoveElementToEnd(array, toMove))

	array2 := []int{1, 2, 2, 2, 2, 1, 4, 3, 2}
	fmt.Println("Order-Preserving:", MoveElementToEndPreservingOrder(array2, toMove))
}
```
