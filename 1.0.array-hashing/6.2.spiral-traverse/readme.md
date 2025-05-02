# Spiral Traversal in Go (Iterative vs Recursive)

This README explains two different methods to traverse a 2D array (matrix) in **spiral order**, implemented in **Go**:

1. Iterative Approach
2. Recursive Approach

We'll walk through both, visually and conceptually.

---

## 📌 Problem Statement

Given a 2D matrix like this:

```
[ [1,  2,  3,  4],
  [5,  6,  7,  8],
  [9, 10, 11, 12] ]
```

The spiral order traversal is:

```
[1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7]
```

---

## 🌀 1. Iterative Spiral Traversal

### ✅ Go Code

```go
func spiralTraverse(matrix [][]int) []int {
    result := []int{}
    if len(matrix) == 0 {
        return result
    }

    top, bottom := 0, len(matrix)-1
    left, right := 0, len(matrix[0])-1

    for top <= bottom && left <= right {
        for col := left; col <= right; col++ {
            result = append(result, matrix[top][col])
        }
        top++

        for row := top; row <= bottom; row++ {
            result = append(result, matrix[row][right])
        }
        right--

        if top <= bottom {
            for col := right; col >= left; col-- {
                result = append(result, matrix[bottom][col])
            }
            bottom--
        }

        if left <= right {
            for row := bottom; row >= top; row-- {
                result = append(result, matrix[row][left])
            }
            left++
        }
    }

    return result
}
```

### 🔍 Step-by-Step Visualization

Initial boundaries:

```
Top = 0, Bottom = 2, Left = 0, Right = 3
```

#### Layer 1:

```
Traverse →   [1, 2, 3, 4]
Traverse ↓   [8, 12]
Traverse ←   [11, 10, 9]
Traverse ↑   [5]
```

Then shrink boundaries:

```
Top = 1, Bottom = 1, Left = 1, Right = 2
```

#### Layer 2:

```
Traverse →   [6, 7]
```

Loop ends.

### ✅ Output:

```
[1 2 3 4 8 12 11 10 9 5 6 7]
```

---

## 🔁 2. Recursive Spiral Traversal

### ✅ Go Code

```go
func spiralTraverse(matrix [][]int) []int {
    result := []int{}
    if len(matrix) == 0 {
        return result
    }
    spiralFill(matrix, 0, len(matrix)-1, 0, len(matrix[0])-1, &result)
    return result
}

func spiralFill(matrix [][]int, startRow, endRow, startCol, endCol int, result *[]int) {
    if startRow > endRow || startCol > endCol {
        return
    }

    for col := startCol; col <= endCol; col++ {
        *result = append(*result, matrix[startRow][col])
    }

    for row := startRow + 1; row <= endRow; row++ {
        *result = append(*result, matrix[row][endCol])
    }

    if startRow < endRow {
        for col := endCol - 1; col >= startCol; col-- {
            *result = append(*result, matrix[endRow][col])
        }
    }

    if startCol < endCol {
        for row := endRow - 1; row > startRow; row-- {
            *result = append(*result, matrix[row][startCol])
        }
    }

    spiralFill(matrix, startRow+1, endRow-1, startCol+1, endCol-1, result)
}
```

### 🔍 Step-by-Step Visualization

#### First Call:

```
spiralFill(matrix, 0, 2, 0, 3)
Traverse → [1 2 3 4]
Traverse ↓ [8 12]
Traverse ← [11 10 9]
Traverse ↑ [5]
```

#### Second Call:

```
spiralFill(matrix, 1, 1, 1, 2)
Traverse → [6 7]
```

#### Third Call:

```
spiralFill(matrix, 2, 0, 2, 1) // base case, ends recursion
```

### ✅ Output:

```
[1 2 3 4 8 12 11 10 9 5 6 7]
```

---

## ⚖️ Iterative vs Recursive

| Feature                | Iterative  | Recursive                 |
| ---------------------- | ---------- | ------------------------- |
| Code Structure         | Uses loops | Uses recursion            |
| Easy to Understand     | ✅          | ✅                         |
| Risk of Stack Overflow | ❌          | ⚠️ (on very large inputs) |
| Good for Visualization | ✅          | ✅                         |
| Performance            | Excellent  | Slightly slower           |

---

## ✅ Summary

Both approaches achieve the same result. Choose:

* **Iterative** if you want performance and control.
* **Recursive** for elegance and layered thinking.

---

Want a diagram of the recursion tree or a Go Playground link?
