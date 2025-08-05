Here's a well-structured `README.md` for your `increasingTriplet` Go function:

---

# Increasing Triplet Subsequence in Go

This Go function checks whether a given slice of integers contains an **increasing subsequence of three elements** (a triplet) in order — i.e., whether there exist indices `i < j < k` such that `nums[i] < nums[j] < nums[k]`.

---

## 🔍 Problem Statement

Given an integer array `nums`, return `true` if there exists a triplet `(i, j, k)` such that:

```
nums[i] < nums[j] < nums[k]` and `i < j < k
```

Otherwise, return `false`.

---

## ✅ Example

### Input:

```go
[]int{1, 2, 3, 4, 5}
```

### Output:

```go
true
```

Because `1 < 2 < 3` is an increasing triplet.

---

### Input:

```go
[]int{5, 4, 3, 2, 1}
```

### Output:

```go
false
```

No increasing triplet exists.

---

## ⚙️ Algorithm Explained

This algorithm uses **two pointers** — `first` and `second` — to track the smallest and second smallest numbers seen so far.

### Pseudocode:

* Initialize `first` and `second` to a large value (`MaxInt`)
* Iterate through the array:

  * If `n <= first`, update `first`
  * Else if `n <= second`, update `second`
  * Else, `n` is greater than both → increasing triplet found

### Code:

```go
func increasingTriplet(nums []int) bool {
    first := int(^uint(0) >> 1)  // MaxInt
    second := int(^uint(0) >> 1) // MaxInt

    for _, n := range nums {
        if n <= first {
            first = n
        } else if n <= second {
            second = n
        } else {
            return true
        }
    }

    return false
}
```

---

## ⏱️ Time and Space Complexity

* **Time Complexity:** `O(n)` — single pass through the array
* **Space Complexity:** `O(1)` — constant extra space

---

## 📦 Usage

```go
func main() {
    fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6})) // Output: true
    fmt.Println(increasingTriplet([]int{20, 100, 10, 12, 5, 13})) // Output: true
    fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1})) // Output: false
}
```

---

## 📌 Notes

* This algorithm does **not** return the actual triplet, just whether one exists.
* It’s optimal for large input arrays due to its `O(n)` time and `O(1)` space behavior.

---

## 🧪 Test Cases

| Input                      | Output  | Reason                     |
| -------------------------- | ------- | -------------------------- |
| `[1, 2, 3, 4, 5]`          | `true`  | Obvious increasing triplet |
| `[5, 4, 3, 2, 1]`          | `false` | Strictly decreasing        |
| `[2, 1, 5, 0, 4, 6]`       | `true`  | Triplet: `0 < 4 < 6`       |
| `[20, 100, 10, 12, 5, 13]` | `true`  | Triplet: `10 < 12 < 13`    |

---

## 🧠 Inspiration

This approach is inspired by the **patience sorting** idea used in algorithms like Longest Increasing Subsequence but optimized specifically for triplets.
