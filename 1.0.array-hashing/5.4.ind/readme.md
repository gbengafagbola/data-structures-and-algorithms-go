Sure! Here's a polished `README.md` file for your Go package:

---

# 📦 numberOfItems

A Go package that provides efficient methods to solve the "Items Between Containers" problem — determining the number of items (`*`) enclosed between container symbols (`|`) within given query ranges of a string.

---

## 📖 Problem Statement

Given a string consisting of `|` and `*` characters, where `|` represents container boundaries and `*` represents items, the goal is to answer multiple queries. Each query provides a `startIndex` and an `endIndex`, and you must determine how many items are enclosed between the outermost pair of containers (`|`) in that range.

---

## ✅ Features

- Three implementations:
  - `numberOfItems`: Optimized using prefix sums and indexed container positions.
  - `numberOfItems2`: Simpler approach iterating through substrings.
  - `numberOfItems3`: Hardcoded for exactly two queries with an alternate approach.

---

## 🚀 Usage

### 📄 Example Input

```go
startIndices := []int32{1, 1}
endIndices := []int32{5, 6}
s := "|**|*|*"
```

### 🧮 Expected Output

```bash
[2, 3]
```

This means:
- Query 1: Between index 1 and 5 => `|**|*` → 2 items between first and second `|`
- Query 2: Between index 1 and 6 => `|**|*|` → 3 items between the first and last `|`

---

## 🛠️ Functions

### `numberOfItems(s string, startIndices, endIndices []int32) []int32`

Efficient solution using prefix sums and pre-indexed `|` characters for faster range queries.

### `numberOfItems2(s string, startIndices, endIndices []int32) []int32`

More straightforward, substring-based approach. Simpler, but less performant for large input.

### `numberOfItems3(s string, startIndices, endIndices []int32) []int32`

Hardcoded logic for exactly 2 queries using ASCII arithmetic. Useful for demo or limited scope use.

---

## 📦 Running the Program

```bash
go run main.go
```

---

## 🧪 Sample Output

```bash
[2 3]
```

---

## 📌 Notes

- Indices in queries are 1-based.
- Only valid `|`-bounded item ranges are considered.
- The main function demonstrates usage with a sample string and two queries.

---

## 📂 File Structure

```
main.go         # Contains all function implementations and sample main()
README.md       # You're reading it!
```

---

## 👨‍💻 Author

**Oluwagbenga Fagbola**
