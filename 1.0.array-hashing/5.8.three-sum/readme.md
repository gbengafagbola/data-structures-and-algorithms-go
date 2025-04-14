Here's a simple and clear `README.md` file for your `ThreeNumberSum` Go implementation:

---

```markdown
# Three Number Sum in Go

This Go program provides a function `ThreeNumberSum` that finds all triplets in an integer array that sum up to a given target value.

## 🧠 Problem Statement

Given an array of integers and a target sum, find all unique triplets in the array that add up to the target sum.

The input array may contain both positive and negative integers. The output triplets must be in ascending order and sorted based on the order of the first element, then the second, and so on.

## 📦 Function Signature

```go
func ThreeNumberSum(array []int, target int) [][]int
```

### Parameters:

- `array []int`: The list of integers to search.
- `target int`: The target sum for the triplets.

### Returns:

- A slice of integer slices (`[][]int`) where each inner slice is a triplet that sums to the target value.

## 🧪 Example

```go
array := []int{12, 3, 1, 2, -6, 5, -8, 6}
target := 0
result := ThreeNumberSum(array, target)

// Output:
// [
//   [-8, 2, 6],
//   [-8, 3, 5],
//   [-6, 1, 5],
//   [-6, 2, 4]
// ]
```

## 🚀 How It Works

1. The input array is first sorted in ascending order.
2. For each element in the array (up to the third last), a two-pointer technique is used to find pairs that sum with the current element to the target.
3. All valid triplets are collected in the result slice.

## 🛠️ How to Run

Make sure you have Go installed. Save the code in a file, for example `main.go`, and then run:

```bash
go run main.go
```

You can test the `ThreeNumberSum` function inside the `main` function.

