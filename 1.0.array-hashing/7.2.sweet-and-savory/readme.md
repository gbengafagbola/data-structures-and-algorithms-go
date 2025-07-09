Here’s a comprehensive `README.md` for your `SweetAndSavory` function:

---

# 🍽️ Sweet and Savory Pair Finder

This Go program implements a function that finds the best combination of one **sweet dish** and one **savory dish** such that:

* The total flavor value is **as close as possible to a given `target`** without **exceeding** it.
* A **sweet dish** is represented by a **negative integer**, while a **savory dish** is represented by a **positive integer**.
* The function ensures one value from each category is selected (`sweet < 0`, `savory > 0`).

---

## 📦 Function Signature

```go
func SweetAndSavory(dishes []int, target int) []int
```

* **Input**:

  * `dishes`: A list of integers representing flavor values of available dishes.
  * `target`: An integer representing the desired combined flavor value.
* **Output**:

  * A slice of two integers: `[sweet, savory]` representing the best sweet-savory pair.
  * Returns an empty list if no such pair exists.
  * Returns `[0, 0]` if `dishes` is empty.

---

## 🚦 Assumptions

* There must be **at least one sweet and one savory** dish to form a valid pair.
* The returned sum must be **less than or equal to** the target.
* The pair should be the **closest to the target** from below or equal.

---

## 🔄 How It Works

1. **Initial Setup**:

   * The input array is sorted to enable two-pointer traversal.
   * Variables `i` and `j` are initialized to the start and end of the array, respectively.

2. **Edge Case Handling**:

   * If `dishes` is empty, the function returns `[0, 0]`.

3. **Two-Pointer Strategy**:

   * Pointer `i` looks for **sweet dishes** (negative values).
   * Pointer `j` looks for **savory dishes** (positive values).
   * The loop runs while `i < j` and ensures that one value is negative (sweet) and the other is positive (savory).

4. **Choosing Best Pair**:

   * For each valid sweet-savory combination, calculate the `sum`.
   * If `sum <= target` and closer to the target than previous best, update the result.

---

## ✅ Example

```go
SweetAndSavory([]int{-3, -1, 4, 5, 10}, 8)
// Output: [-3, 10] because -3 + 10 = 7 is the closest to 8 without exceeding it.
```

---

## 🧪 Edge Cases

```go
SweetAndSavory([]int{}, 10)
// Output: [0, 0]

SweetAndSavory([]int{-5, -2, -1}, 3)
// Output: [] (no savory dishes)

SweetAndSavory([]int{4, 5, 6}, 10)
// Output: [] (no sweet dishes)

SweetAndSavory([]int{-5, 6}, 0)
// Output: [-5, 6] (sum = 1, exceeds target → not valid)
```

---

## 🛠️ Complexity

* **Time**: O(n log n) for sorting + O(n) for traversal ⇒ **O(n log n)**
* **Space**: O(1) extra space (excluding result)

---

## 🔚 Conclusion

This function efficiently solves the "sweet and savory pairing" problem using sorting and a two-pointer technique. It's ideal for finding optimized pairs under constraints like taste balance and target limits.

---

Let me know if you want a version with unit tests or CLI support!
