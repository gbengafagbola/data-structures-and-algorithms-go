# 🧮 Missing Number Finder

## 📌 Overview
Ever had a sequence of numbers and realized one was missing? This Go program helps you find that missing number efficiently! Using mathematical logic and iteration, it quickly determines which number is absent from a given sequence.

## 🚀 How It Works
This program includes two different approaches to solving the missing number problem:

### 1️⃣ Gauss Formula Method
Using the well-known Gauss formula for the sum of the first `n` natural numbers:
```
Sum = n * (n + 1) / 2
```
This method calculates the intended sum and subtracts the actual sum to find the missing number.

### 2️⃣ Iterative Sum Method
Instead of using the formula, this method manually sums up the given numbers and compares them to the expected sum.

## 📜 Code Explanation
```go
// Mathematical approach using Gauss formula
func missingNumber(nums []int) int {
    n := len(nums)
    currentSum := 0
    intendedSum := n * (n + 1) / 2

    for i := 0; i < n; i++ {
        currentSum += i
    }

    return intendedSum - currentSum
}
```
🔹 **Problem:** The loop incorrectly sums up `i` instead of `nums[i]`. This results in incorrect outputs.

```go
// Iterative approach
func missingNumber2(nums []int) int {
    n := len(nums)
    currentSum := 0
    intendedSum := 0

    for i := 0; i < n; i++ {
        currentSum += i
    }

    for i := 0; i < (n+1); i++ {
        intendedSum += i
    }

    return intendedSum - currentSum
}
```
🔹 **Problem:** This method also sums `i` instead of `nums[i]`, leading to an incorrect missing number calculation.

## ✅ Corrected Approach
```go
func missingNumber(nums []int) int {
    n := len(nums)
    sum := 0
    expectedSum := n * (n + 1) / 2

    for _, num := range nums {
        sum += num
    }

    return expectedSum - sum
}
```
Now, the loop correctly sums `nums[i]`, ensuring an accurate result.

## 🔧 How to Run
1. Clone the repository or copy the code.
2. Open a terminal and navigate to the file location.
3. Run the program with:
   ```sh
   go run main.go
   ```
4. The missing number will be printed in the console.

## 🎯 Example Input & Output
```go
nums := []int{1,3,4,5,6}
fmt.Println(missingNumber(nums))
```
📝 **Expected Output:**
```
2
```

## 🎨 Visual Representation
```plaintext
Original Sequence: [1, 3, 4, 5, 6]
Full Sequence:     [1, 2, 3, 4, 5, 6]
Missing Number:    2
```

## 📌 Key Takeaways
- Uses mathematical and iterative approaches to find the missing number.
- Demonstrates efficient problem-solving using Go.
- Shows the importance of correctly iterating over arrays.

[source](https://leetcode.com/problems/missing-number/description/)
