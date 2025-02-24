# 🤖 Judge Circle - Robot Movement Tracker

## 📌 Problem Statement
A robot starts at the **origin (0,0)** and is given a series of movements as a string containing characters:

- **'U'** (Up) → Moves the robot **+1 in the y-axis**
- **'D'** (Down) → Moves the robot **-1 in the y-axis**
- **'R'** (Right) → Moves the robot **+1 in the x-axis**
- **'L'** (Left) → Moves the robot **-1 in the x-axis**

After following all the movements, the function should return `true` if the robot **returns to the origin (0,0)**; otherwise, it returns `false`.

---

## 💻 Code Implementation in Go

```go
package main

func judgeCircle(moves string) bool {
    x, y := 0, 0

    for _, move := range moves {
        switch move {
        case 'U':
            y++
        case 'D':
            y--
        case 'R':
            x++
        case 'L':
            x--
        }
    }

    return x == 0 && y == 0
}
```

---

## 🔍 How It Works
1. Initialize the robot at `(0,0)`.
2. Iterate through each move in the string:
   - **'U'** increases `y`
   - **'D'** decreases `y`
   - **'R'** increases `x`
   - **'L'** decreases `x`
3. After all moves are processed, check if `x == 0 && y == 0`.
   - If **true**, the robot is back at the starting point → return `true` ✅.
   - If **false**, the robot is not at the origin → return `false` ❌.

---

## 📌 Example Walkthroughs

### **Example 1**
```go
moves := "UD"
fmt.Println(judgeCircle(moves)) // Output: true
```
🔹 Explanation:
- 'U' → (0,1)
- 'D' → (0,0) ✅

### **Example 2**
```go
moves := "LLRR"
fmt.Println(judgeCircle(moves)) // Output: true
```
🔹 Explanation:
- 'L' → (-1,0)
- 'L' → (-2,0)
- 'R' → (-1,0)
- 'R' → (0,0) ✅

### **Example 3**
```go
moves := "UDLRUD"
fmt.Println(judgeCircle(moves)) // Output: false
```
🔹 Explanation:
- 'U' → (0,1)
- 'D' → (0,0)
- 'L' → (-1,0)
- 'R' → (0,0)
- 'U' → (0,1)
- 'D' → (0,0) ✅

---

## 🔥 Why This Works Efficiently
- **Time Complexity:** `O(n)`, since we iterate through the moves once.
- **Space Complexity:** `O(1)`, as we only use two integer variables (`x` and `y`).

### 🚀 Summary
✔ Uses **simple coordinate updates** to track movement.
✔ Checks if the robot **returns to the origin**.
✔ Efficient and **easy to understand**.

Let me know if you have any questions! 🤖🔥

