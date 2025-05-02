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

Perfect. Let's walk through the **spiral traversal algorithm step by step**, using a **visual example** and explaining how the bounds (`top`, `bottom`, `left`, `right`) define and shrink the spiral layer by layer.

---

### 📊 Example Matrix

Let’s use this 3x4 matrix:

```
matrix = [
  [1,  2,  3,  4],
  [5,  6,  7,  8],
  [9, 10, 11, 12]
]
```

At the start:

* `top = 0` (first row)
* `bottom = 2` (last row, index 2)
* `left = 0` (first column)
* `right = 3` (last column, index 3)

---

### 🔁 First Spiral Layer

**1. Traverse → right** along `top = 0` from `left` to `right`:

```
[1, 2, 3, 4]
 ↑  ↑  ↑  ↑
```

→ Append: `1, 2, 3, 4`
→ Then: `top++` → `top = 1`

**2. Traverse ↓ down** along `right = 3` from `top` to `bottom`:

```
[       8]
[       12]
```

→ Append: `8, 12`
→ Then: `right--` → `right = 2`

**3. Traverse ← left** along `bottom = 2` from `right` to `left`, only if `top <= bottom`:

```
[11, 10, 9]
   ↑   ↑  ↑
```

→ Append: `11, 10, 9`
→ Then: `bottom--` → `bottom = 1`

**4. Traverse ↑ up** along `left = 0` from `bottom` to `top`, only if `left <= right`:

```
[5]
```

→ Append: `5`
→ Then: `left++` → `left = 1`

---

### 🌀 Second Spiral Layer

Updated bounds:

* `top = 1`
* `bottom = 1`
* `left = 1`
* `right = 2`

**5. Traverse → right** again along `top = 1` from `left` to `right`:

```
[6, 7]
 ↑  ↑
```

→ Append: `6, 7`
→ Then: `top++` → `top = 2`

Now `top > bottom` → loop ends.

---

### ✅ Final Output

Result slice:

```
[1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7]
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
