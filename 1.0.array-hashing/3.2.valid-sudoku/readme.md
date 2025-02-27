# Valid Sudoku Explanation

## What Does This Function Do?
It checks whether a given **9x9 Sudoku board** is valid. A Sudoku board is valid if:
1. **Each row** has unique numbers (1-9, no duplicates).
2. **Each column** has unique numbers.
3. **Each 3x3 small box** has unique numbers.

The board may contain empty cells (represented as `.`), which are ignored.

---

## Step-by-Step Explanation

### Step 1: Check Each Row
- We go **row by row** and check if there are any duplicate numbers.
- We use a **map (`seen`)** to track numbers we have already encountered.

**Example:**
```
Row 0: [5, 3, ., ., 7, ., ., ., .] ✅ (No duplicates)
Row 1: [6, ., ., 1, 9, 5, ., ., .] ✅
Row 2: [., 9, 8, ., ., ., ., 6, .] ✅
```
If a duplicate is found, return `false`.

---

### Step 2: Check Each Column
- We go **column by column** and check for duplicates using a similar map.

**Example (Column 0):**
```
Column 0: [5, 6, ., 8, 4, 7, ., ., .] ✅ (No duplicates)
```
If a duplicate is found, return `false`.

---

### Step 3: Check Each 3x3 Box
- Sudoku is divided into **nine** `3x3` boxes. We ensure each box contains unique numbers.

**Example:**
```
Box 0 (Top-left 3x3): ✅
[5, 3, .]
[6, ., .]
[., 9, 8]

Box 1 (Top-middle 3x3): ✅
[., 7, .]
[1, 9, 5]
[., ., .]
```
If a duplicate is found, return `false`.

---

## Visual Representation

### Example Sudoku Board
```
[
  ['5','3','.','.','7','.','.','.','.'],
  ['6','.','.','1','9','5','.','.','.'],
  ['.','9','8','.','.','.','.','6','.'],
  ['8','.','.','.','6','.','.','.','3'],
  ['4','.','.','8','.','3','.','.','1'],
  ['7','.','.','.','2','.','.','.','6'],
  ['.','6','.','.','.','.','2','8','.'],
  ['.','.','.','4','1','9','.','.','5'],
  ['.','.','.','.','8','.','.','7','9']
]
```
✅ **This board is valid!** The function will return `true`.

---

## How Does the Code Work?
```go
for row := 0; row < 9; row++ { // Check rows
    seen := make(map[byte]bool)
    for i := 0; i < 9; i++ {
        if board[row][i] == '.' { continue }
        if seen[board[row][i]] { return false } // Duplicate found
        seen[board[row][i]] = true
    }
}
```
✔️ Goes through **each row**, checking for duplicates.

```go
for col := 0; col < 9; col++ { // Check columns
    seen := make(map[byte]bool)
    for i := 0; i < 9; i++ {
        if board[i][col] == '.' { continue }
        if seen[board[i][col]] { return false }
        seen[board[i][col]] = true
    }
}
```
✔️ Goes through **each column**, checking for duplicates.

```go
for square := 0; square < 9; square++ { // Check 3x3 boxes
    seen := make(map[byte]bool)
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            row := (square / 3) * 3 + i
            col := (square % 3) * 3 + j
            if board[row][col] == '.' { continue }
            if seen[board[row][col]] { return false }
            seen[board[row][col]] = true
        }
    }
}
```
✔️ Iterates over each **3x3 box**, checking for duplicates.

---

## Final Thoughts
✅ **Time Complexity:** O(9²) = O(81) ≈ O(1) (since it's always a 9x9 grid)  
✅ **Space Complexity:** O(9) per row/column/box check = O(1)  

This is an **efficient** and **simple** way to validate a Sudoku board. 🚀

