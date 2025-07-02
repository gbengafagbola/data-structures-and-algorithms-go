# 🪑 BestSeat – Find the Most Optimal Seat in a Row

This project provides a Go implementation to solve a common interview-style problem: **"Given a row of seats, find the best empty seat such that you maximize your distance from the closest person."**

---

## 📖 Problem Description

You’re given an array of integers:

* `1` represents an **occupied seat**.
* `0` represents an **empty seat**.

You must find the **index of the empty seat** that lies **in the middle of the longest contiguous group of zeros**. This seat ensures **maximum distance from others**.

### ❗ Important Constraints:

* You can only sit in the **middle** of the longest sequence of empty seats.
* Seats at the very beginning or end of the row (i.e., edge seating) are ignored for simplicity unless otherwise stated.
* If multiple groups have the same length, the **first** one is selected.
* If there are no available seats (`0`s), return `-1`.

---

## 🧠 Algorithm Breakdown – Line-by-Line Explanation

Here is the working function again:

```go
func BestSeat(seats []int) int {
    result := -1         // Default result if no seat is found
    diff := -1           // Keeps track of the longest sequence length

    for i := 1; i < len(seats)-1; {
        j := i

        // Check if current position starts a block of zeros
        if seats[i] == 0 && seats[j] == 0 {
            // Expand 'j' to find the end of the zero block
            for j < len(seats) && seats[j] == 0 {
                j++
            }

            // Calculate middle seat of zero block
            mid := (i + j - 1) / 2
            difference := j - i

            // Update best seat if current block is longer
            if difference > diff {
                diff = difference
                result = mid
            }

            i = j // Move 'i' to the end of the current zero block
        } else {
            i++ // Move to the next index if current isn't zero
        }
    }

    return result
}
```

---

### 🔍 Detailed Line-by-Line Insight

| Line                                  | Explanation                                                         |
| ------------------------------------- | ------------------------------------------------------------------- |
| `result := -1`                        | Initializes the seat index to return; `-1` means no seat was found. |
| `diff := -1`                          | Stores the length of the longest streak of `0`s seen so far.        |
| `for i := 1; i < len(seats)-1;`       | Starts scanning from index 1 to `len(seats)-2` (ignores edges).     |
| `j := i`                              | `j` is a helper pointer used to scan ahead in the array.            |
| `if seats[i] == 0 && seats[j] == 0`   | Confirms this is the start of a zero streak.                        |
| `for j < len(seats) && seats[j] == 0` | Finds the end of the zero block.                                    |
| `mid := (i + j - 1) / 2`              | Calculates the middle index of the current zero block.              |
| `difference := j - i`                 | Calculates how long the block of empty seats is.                    |
| `if difference > diff`                | Compares it with the previously longest block.                      |
| `diff = difference` & `result = mid`  | Updates the best result.                                            |
| `i = j`                               | Skips all the `0`s already processed.                               |
| `else { i++ }`                        | If current seat is not zero, just move forward.                     |
| `return result`                       | Returns the best seat index found.                                  |

---

## 🧪 Example

### Input:

```go
seats := []int{1, 0, 0, 0, 1, 0, 1}
```

### Process:

* Sequence of `0, 0, 0` from index `1` to `3`.
* Middle is `(1 + 3) / 2 = 2`.
* That's the longest zero block.
* A later single `0` at index `5` is shorter.

### Output:

```
Best seat index: 2
```

---

## 🧮 Time & Space Complexity

| Metric | Value                                                           |
| ------ | --------------------------------------------------------------- |
| Time   | **O(n)** – Every element is scanned at most twice (`i` and `j`) |
| Space  | **O(1)** – No extra space used                                  |

No nested iteration occurs over the same elements, so complexity remains linear.

---

## ⚠️ Edge Cases & Considerations

| Case                                         | Handling                                     |
| -------------------------------------------- | -------------------------------------------- |
| All seats are occupied                       | Returns `-1`                                 |
| No block of 0s longer than 1                 | Still returns the middle of that block       |
| Edge seats (start or end of array) are empty | Ignored by logic (can be modified if needed) |
| Multiple equal-length blocks                 | The **first one** is returned                |

---

## 🔧 How to Run

1. **Create a file** named `main.go`.
2. Paste in the following code:

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

func main() {
    seats := []int{1, 0, 0, 0, 1, 0, 1}
    fmt.Println("Best seat index:", BestSeat(seats))
}
```

3. Run it:

```bash
go run main.go
```

---

## 📌 Customization Ideas

* Allow seating at the **edges** if empty.
* Accept real-time input from the user.
* Return **multiple seat options** if multiple sequences are tied.
* Track the **distance to nearest person**.

---

## 🧪 Optional Tests

Here are some sample test cases you can run:

| Input                   | Expected Output | Explanation                       |
| ----------------------- | --------------- | --------------------------------- |
| `[1, 0, 0, 0, 1]`       | `2`             | Longest zero block: `1–3`         |
| `[1, 0, 1, 0, 0, 0, 1]` | `4`             | Longest block: `3–5`, middle is 4 |
| `[1, 1, 1]`             | `-1`            | No empty seats                    |
| `[1, 0, 1]`             | `1`             | Only one zero                     |

---

## 🏁 Conclusion

This Go implementation provides a clean, efficient, and readable solution to locating the best possible seat in a row, based on maximizing distance from other occupied seats. It's built to run in linear time with constant space, while allowing for easy expansion and customization.