Here’s a clean and informative `README.md` that explains both versions of your `IsMonotonic` function — with visuals, complexity breakdown, and code comments. You can drop this straight into your repo 👇

---

```markdown
# 🧮 Monotonic Array Checker in Go

This repository contains two Go implementations of a function to check whether a given array is **monotonic** — that is, always increasing, always decreasing, or remaining the same.

---

## ✅ What is a Monotonic Array?

An array is **monotonic** if:

- It is entirely **non-increasing**  
  _e.g._ `[9, 7, 7, 5, 3]`

**OR**

- It is entirely **non-decreasing**  
  _e.g._ `[1, 2, 2, 3, 4]`

---

## 🚀 Approaches Compared

We implemented and compared two methods:

| Version | Method | Time | Space | Early Exit | Readability | Handles Duplicates |
|--------|--------|------|-------|------------|-------------|--------------------|
| 1️⃣ | Count-based | O(n) | O(1) | ❌ | 🤏 Verbose | ✅ |
| 2️⃣ | Direction-based (preferred) | O(n) | O(1) | ✅ | ✅ Clean | ✅ |

---

## 🧪 Example Inputs & Visuals

| Input | Type | Monotonic? |
|-------|------|------------|
| `[-1, -5, -10, -1100, -1100, -1101, -1102, -9001]` | Decreasing | ✅ Yes |
| `[1, 2, 2, 3, 4]` | Increasing | ✅ Yes |
| `[1, 3, 2]` | Mixed | ❌ No |

---

## 🧠 1. Count-Based Approach

### Code:

```go
func IsMonotonic(array []int) bool {
    if len(array) <= 2 {
        return true
    }

    count := 0
    equals := 0

    for i := 0; i < len(array)-1; i++ {
        j := i + 1
        if array[i] < array[j] {
            count--
        } else if array[i] > array[j] {
            count++
        } else {
            equals++
        }
    }

    if count < 0 {
        count = -count
    }

    return count + equals == len(array)-1
}
```

### 📉 Visual Representation

```
[9, 7, 5, 3]
  ↘  ↘  ↘  → Decreasing trend → count++
```

---

## ⚡ 2. Direction-Based Approach (Recommended)

### Code:

```go
func IsMonotonic(array []int) bool {
    if len(array) <= 2 {
        return true
    }

    direction := array[1] - array[0]

    for i := 2; i < len(array); i++ {
        if direction == 0 {
            direction = array[i] - array[i-1]
            continue
        }

        if breaksDirection(direction, array[i-1], array[i]) {
            return false
        }
    }

    return true
}

func breaksDirection(direction int, prev int, curr int) bool {
    diff := curr - prev
    return direction > 0 && diff < 0 || direction < 0 && diff > 0
}
```

### 📊 Visual Representation

```
Direction = array[1] - array[0] = 2 - 1 = +1

[1, 2, 2, 3]
    ↗   →   ↗ → Still increasing → ✅
```

### 💥 Early Exit:
If the trend breaks midway, this function returns early instead of scanning the whole array.

---

## 🧾 Conclusion

While both approaches work with the same complexity:

- ✅ The **direction-based method** is more elegant and efficient in practice.
- ✅ Handles all edge cases cleanly.
- ✅ Easier to read, debug, and extend.
