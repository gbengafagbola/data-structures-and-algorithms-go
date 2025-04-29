# Monotonic Array Checker in Go

This project provides multiple implementations in Go for checking if an array is **monotonic**.

A **monotonic array** is one where the elements are either entirely non-increasing or non-decreasing.

## 📂 Overview of Functions

### 1. `IsMonotonic(array []int) bool`

Checks monotonicity by assuming the array is both non-increasing and non-decreasing, then invalidating each possibility as it finds violations.

**Logic:**
- Start with both `isNonIncreasing` and `isNonDecreasing` set to `true`.
- Iterate through the array:
  - If any element is greater than its predecessor → not non-increasing.
  - If any element is less than its predecessor → not non-decreasing.
- If either condition remains `true`, the array is monotonic.

**⚠ Bug:** The loop ends at `len(array)-1`, skipping the last element. It should go up to `len(array)`.

---

### 2. `IsMonotonic2(array []int) bool`

Determines monotonicity by detecting the direction (increasing or decreasing) based on the first non-equal pair.

**Logic:**
- Calculate initial `direction = array[1] - array[0]`.
- Loop through the rest of the array:
  - If direction is `0`, update it until it's non-zero.
  - For non-zero direction, check if each subsequent pair breaks it using `breaksDirection`.

---

### 3. `IsMonotonic3(array []int) bool`

This is identical to `IsMonotonic2` and serves as a duplicate for reference or testing purposes.

---

### 4. `IsMonotonic4(array []int) bool`

Counts the number of increases, decreases, and equal pairs to assess consistency.

**Logic:**
- Loop through array pairs:
  - `count++` if decreasing
  - `count--` if increasing
  - `equals++` if equal
- Normalize `count` to positive.
- If `count + equals == len(array) - 1`, then array is monotonic.

---

### 🧰 Helper Function

#### `breaksDirection(direction int, prev int, curr int) bool`

Used by `IsMonotonic2` and `IsMonotonic3` to determine if a value transition violates the identified direction.

- If direction is positive, any drop (curr < prev) breaks it.
- If direction is negative, any rise (curr > prev) breaks it.

---

## ✅ Example

```go
arr := []int{1, 2, 2, 3}

fmt.Println(IsMonotonic(arr))   // true
fmt.Println(IsMonotonic2(arr))  // true
fmt.Println(IsMonotonic3(arr))  // true
fmt.Println(IsMonotonic4(arr))  // true
