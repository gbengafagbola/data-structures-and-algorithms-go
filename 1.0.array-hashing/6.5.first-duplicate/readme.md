# 📦 FirstDuplicateValue – Go Implementation

This Go package provides two implementations of the `FirstDuplicateValue` function, which finds the first integer that appears more than once in a given array. The second occurrence must have the **minimum index** compared to other duplicates.

---

## 🚀 Problem Statement

Given an array of integers, return the **first integer** that appears more than once when traversing the array from left to right.

### ✅ Constraints:

* The array contains integers.
* The integers are all in the range `1` to `n` where `n == len(array)` (for the second approach).
* You are allowed to mutate the array (for the second approach).

### 🔁 Example:

```go
input := []int{2, 1, 5, 2, 3, 3, 4}
FirstDuplicateValue(input) // Output: 2
```

---

## 📘 Usage

```bash
go run main.go
```

You can import or copy the `FirstDuplicateValue` and `FirstDuplicateValue2` functions into your project as needed.

---

## 🧠 Breakdown of Approaches

### ✨ Approach 1: Using a Hash Map

```go
func FirstDuplicateValue(array []int) int {
	n := make(map[int]bool)

	for _, value := range array {
		if _, exists := n[value]; exists {
			return value
		}
		n[value] = true
	}
	return -1
}
```

#### 🔍 How it works:

* Create a map to keep track of seen values.
* Iterate through the array.
* If a value is already in the map, return it.
* If not, store it in the map.
* Return `-1` if no duplicates are found.

#### 📈 Time & Space Complexity:

* **Time:** `O(n)` – each element is checked once.
* **Space:** `O(n)` – for the map used to track seen elements.

---

### ⚡ Approach 2: Using the Array Itself (Optimized Space)

```go
func FirstDuplicateValue2(array []int) int {
	for i := 0; i < len(array); i++ {
		index := abs(array[i]) - 1
		if array[index] < 0 {
			return abs(array[i])
		}
		array[index] = -array[index]
	}
	return -1
}
```

#### 🔍 How it works:

* Leverages the array itself to track seen values.
* For each element, treat its absolute value as an index.
* If the value at that index is negative, it's a duplicate.
* Otherwise, negate the value at that index to mark it as seen.

#### 💡 Why it works:

* All values are within the range `1` to `n`, which ensures valid index access.
* Negation serves as a marker without needing extra space.

#### 📈 Time & Space Complexity:

* **Time:** `O(n)` – iterate through the array once.
* **Space:** `O(1)` – constant space; mutation is done in-place.

#### ⚠️ Caveat:

* This approach **mutates the input array**. If mutation is not allowed, use Approach 1.

---

## 🧪 Utility Function

```go
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

A simple helper function to return the absolute value of an integer.

---

## 📌 Summary Comparison

| Feature          | Approach 1 (Hash Map)          | Approach 2 (In-place)     |
| ---------------- | ------------------------------ | ------------------------- |
| Time Complexity  | O(n)                           | O(n)                      |
| Space Complexity | O(n)                           | O(1)                      |
| Mutates Input    | ❌                              | ✅                         |
| Simplicity       | ✅ Easy to read                 | ⚠️ Slightly trickier      |
| Best Use Case    | When input must not be mutated | When optimizing for space |
