Sure! Here's the **updated README** with:

* The original two solutions
* ✅ A corrected "Safe Placement" version that passes all edge cases
* ⚠️ A clear explanation of why the `threeZeroes` logic fails on some valid inputs
* 🧪 Test case illustrations

---

# 🌼 Can Place Flowers – Go Solutions

This repository provides two different Go implementations for solving the **"Can Place Flowers"** problem, a popular algorithmic challenge.

---

## 🧩 Problem Statement

You're given a slice `flowerbed` of integers where:

* `0` means an empty plot.
* `1` means a plot already has a flower.

You are also given an integer `n`, the number of new flowers you need to plant.
**Rule:** No two flowers can be planted in adjacent plots.

Your goal is to determine whether it's possible to plant `n` new flowers without violating the no-adjacent rule.

---

## ✅ Solution 1: Safe Placement with Boundary Checks

### 🔍 Logic

We scan the flowerbed and check each plot:

* Is it empty (`0`)?
* Is its **left neighbor** also empty or at the start?
* Is its **right neighbor** also empty or at the end?

If all conditions are satisfied:

* We **simulate planting** a flower by setting the current plot to `1`.
* We increment a counter.

If we can plant **at least `n`** flowers this way, we return `true`.

### 💻 Code

```go
func canPlaceFlowers(flowerbed []int, n int) bool {
    numberOfConsecutiveThreeZero := 0

    for i := 0; i < len(flowerbed); i++ {
        if flowerbed[i] == 0 &&
           (i == 0 || flowerbed[i-1] == 0) &&
           (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {

            flowerbed[i] = 1 // Simulate planting
            numberOfConsecutiveThreeZero++
        }
    }

    return numberOfConsecutiveThreeZero >= n
}
```

### 🧠 Complexity

* **Time:** O(n)
* **Space:** O(1)

### ✅ Example

```go
Input:  flowerbed = []int{0, 0, 1, 0, 1}, n = 1  
Output: true
```

Explanation: You can plant a flower at index 0, resulting in `[1, 0, 1, 0, 1]`.

---

## ⚠️ Solution 2: Counting `threeZeroes` (Consecutive Zeros Approach)

### 🔍 Logic

This strategy attempts to place flowers by counting **three consecutive zeroes** (`000`), under the assumption that each block of three allows one flower in the middle.

After planting (virtually), we reset the count to 1 to allow overlapping sequences like `0000`.

### 💻 Code

```go
func canPlaceFlowers(flowerbed []int, n int) bool {
    numberOfConsecutiveThreeZero := 0
    threeZeroes := 0

    for i := 0; i < len(flowerbed); i++ {
        if flowerbed[i] == 0 {
            threeZeroes++
        } else {
            threeZeroes = 0
        }

        if threeZeroes == 3 {
            numberOfConsecutiveThreeZero++
            threeZeroes = 1
        }
    }

    return numberOfConsecutiveThreeZero >= n
}
```

### 🧠 Complexity

* **Time:** O(n)
* **Space:** O(1)

---

## ❌ Known Limitation of `threeZeroes` Strategy

This approach **fails** for edge cases where planting is possible but the pattern isn't `"000"`.

### ❗ Example Failure Case

```go
Input:  flowerbed = []int{0, 0, 1, 0, 1}, n = 1  
Output (Expected): true  
Output (threeZeroes logic): false ❌
```

Explanation:

* A flower can be safely planted at index 0.
* But there are no `"000"` sequences, so the logic misses it.

✅ Use the **Safe Placement (Solution 1)** approach for complete correctness.

---

## 📌 Summary Comparison

| Feature                       | Safe Placement (Solution 1) | threeZeroes (Solution 2) |
| ----------------------------- | --------------------------- | ------------------------ |
| Handles edge cases            | ✅ Yes                       | ❌ No                     |
| Simulates actual planting     | ✅ Yes                       | ❌ No                     |
| Supports overlapping patterns | ✅ Yes                       | ✅ Yes (with reset)       |
| Easy to understand            | ✅ Clean & straightforward   | 🔄 Slightly abstract     |
| Mutates the flowerbed         | ✅ Yes (for simulation)      | ❌ No                     |

---

## ✅ To Run

1. Save the solution in a `.go` file.
2. Run:

```bash
go run main.go
```

---

## 🧪 Additional Test Cases

| Flowerbed         | n | Expected | Reason                           |
| ----------------- | - | -------- | -------------------------------- |
| `[0, 0, 1, 0, 1]` | 1 | `true`   | Can plant at index 0             |
| `[1, 0, 0, 0, 1]` | 1 | `true`   | Can plant at index 2             |
| `[1, 0, 0, 0, 1]` | 2 | `false`  | Only 1 flower can be planted     |
| `[0, 0, 0, 0, 0]` | 2 | `true`   | Plant at index 1 and 3           |
| `[0]`             | 1 | `true`   | Single empty spot, safe to plant |

---

Let me know if you'd like to add diagrams or unit test scaffolds for this!
