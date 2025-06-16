# 📦 Array of Products – Go Implementations

## 🧠 Problem Statement

Given an array of integers, return a new array such that each element at index `i` is the **product of all numbers in the original array except the one at `i`**, and **without using division**.

### 🔍 Example

```go
Input:  [1, 2, 3, 4]
Output: [24, 12, 8, 6]
// Explanation:
// result[0] = 2 * 3 * 4 = 24
// result[1] = 1 * 3 * 4 = 12
// result[2] = 1 * 2 * 4 = 8
// result[3] = 1 * 2 * 3 = 6
```

---

## ✅ Constraints

* No division allowed.
* Time complexity should be **O(n)** for the optimal solution.

---

## ✨ Solutions Breakdown

---

### ⚡️ `ArrayOfProducts` – Optimal 2-Pass Approach

```go
func ArrayOfProducts(array []int) []int {
	n := len(array)
	result := make([]int, n)

	// First pass: prefix products
	prefix := 1
	for i := 0; i < n; i++ {
		result[i] = prefix
		prefix *= array[i]
	}

	// Second pass: suffix products
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= array[i]
	}

	return result
}
```

### 👇 Visual Explanation

**Step 1: Prefix pass (left-to-right)**
Build product of all elements **before** index `i`.

| Index | Prefix Product | result\[i] |
| ----- | -------------- | ---------- |
| 0     | 1              | 1          |
| 1     | 1×1 = 1        | 1          |
| 2     | 1×2 = 2        | 2          |
| 3     | 2×3 = 6        | 6          |

**Step 2: Suffix pass (right-to-left)**
Multiply with product of all elements **after** index `i`.

| Index | Suffix Product | result\[i] |
| ----- | -------------- | ---------- |
| 3     | 1              | 6×1 = 6    |
| 2     | 1×4 = 4        | 2×4 = 8    |
| 1     | 4×3 = 12       | 1×12 = 12  |
| 0     | 12×2 = 24      | 1×24 = 24  |

**✅ Final Output:** `[24, 12, 8, 6]`

**🕒 Time Complexity:** O(n)
**🧠 Space Complexity:** O(1) (excluding output)

---

### ✨ `ArrayOfProducts2` – Verbose with Explicit Prefix/Suffix Arrays

```go
func ArrayOfProducts2(array []int) []int {
	n := len(array)
	result := make([]int, n)
	leftToRightProducts := make([]int, n)
	rightToLeftProducts := make([]int, n)

	prefix := 1
	for i := 0; i < n; i++ {
		leftToRightProducts[i] = prefix
		prefix *= array[i]
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		rightToLeftProducts[i] = suffix
		suffix *= array[i]
	}

	for i := 0; i < n; i++ {
		result[i] = leftToRightProducts[i] * rightToLeftProducts[i]
	}

	return result
}
```

**💡 Benefit:** Easier to understand since prefix and suffix arrays are explicitly separated.
**🕒 Time Complexity:** O(n)
**🧠 Space Complexity:** O(n)

---

### 🐢 `ArrayOfProducts3` – Slice Manipulation (Less Efficient)

```go
func ArrayOfProducts3(array []int) []int {
	result := []int{}

	for i := 0; i < len(array); i++ {
		currentProduct := 1
		x := []int{}

		if i != len(array)-1 {
			x = append(x, array[i+1:]...)
		}
		if i != 0 {
			x = append(x, array[:i]...)
		}

		for _, y := range x {
			currentProduct *= y
		}

		result = append(result, currentProduct)
	}
	return result
}
```

**🔍 How It Works:**

* For every index, create a new slice omitting `array[i]`.
* Multiply all elements in the new slice.

**🛑 Drawbacks:**

* Creating new slices → O(n²) time and high memory use.

**🕒 Time Complexity:** O(n²)
**🧠 Space Complexity:** O(n²) worst case

---

### 🐌 `ArrayOfProducts4` – Brute Force with Manual Multiplication

```go
func ArrayOfProducts4(array []int) []int {
	n := len(array)
	products := make([]int, n)

	for i := range array {
		runningProduct := 1
		for j := range array {
			if i != j {
				runningProduct *= array[j]
			}
		}
		products[i] = runningProduct
	}

	return products
}
```

**🛠 Strategy:**

* Multiply all elements except at index `i` using two nested loops.

**🕒 Time Complexity:** O(n²)
**🧠 Space Complexity:** O(n)

---

## 📊 Comparison Table

| Function           | Time Complexity | Space Complexity | Approach              | Notes                |
| ------------------ | --------------- | ---------------- | --------------------- | -------------------- |
| `ArrayOfProducts`  | O(n)            | O(1)             | Prefix + Suffix       | Most efficient       |
| `ArrayOfProducts2` | O(n)            | O(n)             | Verbose Prefix/Suffix | Easier to read       |
| `ArrayOfProducts3` | O(n²)           | O(n²)            | Slice manipulation    | Not optimal          |
| `ArrayOfProducts4` | O(n²)           | O(n)             | Brute force           | Simplest but slowest |

---

## 🔚 Summary

The `ArrayOfProducts` algorithm is a classic question testing your ability to:

* Avoid division
* Use prefix/suffix accumulation
* Optimize time and space

If you're preparing for interviews, **mastering the prefix-suffix pattern** is essential for efficient array manipulations.

---

Would you like this README saved as a file or formatted for GitHub?
