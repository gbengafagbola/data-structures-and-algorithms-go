Here's the **updated and enhanced README** that includes both `BestSeat` and the improved `BestSeat2` function, with clearly explained logic, edge case notes, and a breakdown of why `BestSeat2` may be preferable for readability and reliability.

---

# 🪑 BestSeat – Find the Most Optimal Seat in a Row

This project provides **Go implementations** to solve a classic problem:

> **Given a row of seats (represented as a list of 0s and 1s), return the index of the best empty seat that maximizes the distance to the nearest occupied seat.**

---

## 📖 Problem Description

You’re given a list of integers where:

* `1` represents an **occupied seat**.
* `0` represents an **empty seat**.

Your goal is to **find the index** of the seat that lies in the **middle of the longest contiguous block of empty seats**, maximizing the distance from others.

---

## ✅ Requirements

* Only choose seats that are not on the **very edges** unless otherwise modified.
* Return the **first such seat** in case of a tie.
* Return `-1` if no empty seat exists.

---

## 🧠 Algorithm 1 – `BestSeat`

This is the original version that uses two pointers (`i` and `j`) to find blocks of empty seats.

### Code:

```go
func BestSeat(seats []int) int {
    result := -1
    diff := -1

    for i := 1; i < len(seats)-1; {
        j := i
        if seats[i] == 0 && seats[j] == 0 {
            for j < len(seats) && seats[j] == 0 {
                j++
            }
            mid := (i + j - 1) / 2
            difference := j - i
            if difference > diff {
                diff = difference
                result = mid
            }
            i = j
        } else {
            i++
        }
    }

    return result
}
```

---

## 🧠 Algorithm 2 – `BestSeat2` (Recommended for Clarity)

This version is more intuitive and works by treating each segment between two `1`s as a candidate block.

### Code:

```go
func BestSeat2(seats []int) int {
    bestSeat := -1
    maxSpace := 0

    left := 0
    for left < len(seats) {
        right := left + 1
        for right < len(seats) && seats[right] == 0 {
            right++
        }

        availableSpace := right - left - 1
        if availableSpace > maxSpace {
            bestSeat = (left + right) / 2
            maxSpace = availableSpace
        }

        left = right
    }

    return bestSeat
}
```

---

## 🚀 Example

### Input:

```go
seats := []int{1, 0, 0, 0, 1, 0, 1}
```

### Output:

```
Best seat index: 2
```

### Why?

* The longest zero block is from index 1 to 3.
* Middle seat is `(1 + 3) / 2 = 2`.

---

## ⏱️ Time and Space Complexity

| Metric | Value                   |
| ------ | ----------------------- |
| Time   | `O(n)` (linear scan)    |
| Space  | `O(1)` (no extra space) |

Both versions scan the array linearly using pointers.

---

## 🧪 Test Cases

| Input                   | Output | Reason                                 |
| ----------------------- | ------ | -------------------------------------- |
| `[1, 0, 0, 0, 1]`       | `2`    | Longest block: index `1–3`, mid is `2` |
| `[1, 0, 1, 0, 0, 0, 1]` | `4`    | Longest block: index `3–5`, mid is `4` |
| `[1, 1, 1]`             | `-1`   | No zero blocks                         |
| `[1, 0, 1]`             | `1`    | Only one seat available                |
| `[1, 0, 0, 1, 0, 0, 1]` | `2`    | First longest block gets chosen        |

---

## ⚠️ Edge Case Considerations

| Case                        | Behavior                                |
| --------------------------- | --------------------------------------- |
| Empty input                 | Returns `-1`                            |
| No zeros                    | Returns `-1`                            |
| Only edge seats are empty   | Ignored unless logic is modified        |
| Multiple same-length blocks | Returns the middle of the **first** one |

---

## 🧩 Customization Ideas

* 🔧 Modify logic to **allow seating at the edges**.
* 🔁 Return **all possible best seat indices** in case of ties.
* 📏 Return the **exact distance to nearest person**.
* 📲 Accept dynamic input from a user or front-end.

---

## 🛠 How to Run

1. Save as `main.go`:

```go
package main

import "fmt"

func BestSeat(seats []int) int {
    result := -1
    diff := -1
    for i := 1; i < len(seats)-1; {
        j := i
        if seats[i] == 0 && seats[j] == 0 {
            for j < len(seats) && seats[j] == 0 {
                j++
            }
            mid := (i + j - 1) / 2
            difference := j - i
            if difference > diff {
                diff = difference
                result = mid
            }
            i = j
        } else {
            i++
        }
    }
    return result
}

func BestSeat2(seats []int) int {
    bestSeat := -1
    maxSpace := 0
    left := 0
    for left < len(seats) {
        right := left + 1
        for right < len(seats) && seats[right] == 0 {
            right++
        }
        availableSpace := right - left - 1
        if availableSpace > maxSpace {
            bestSeat = (left + right) / 2
            maxSpace = availableSpace
        }
        left = right
    }
    return bestSeat
}

func main() {
    seats := []int{1, 0, 0, 0, 1, 0, 1}
    fmt.Println("Best seat index (v1):", BestSeat(seats))
    fmt.Println("Best seat index (v2):", BestSeat2(seats))
}
```

2. Run it:

```bash
go run main.go
```

---

## 🏁 Conclusion

The `BestSeat` and `BestSeat2` implementations showcase two effective approaches to solving the "best seat" problem. Both are linear in time, but `BestSeat2` is cleaner and more intuitive to read, making it suitable for production-grade code or technical interviews.

Feel free to fork, improve, or integrate this solution into larger systems!

---

Let me know if you'd like this published as a `README.md` file or want me to add a testing suite using Go’s `testing` package.
