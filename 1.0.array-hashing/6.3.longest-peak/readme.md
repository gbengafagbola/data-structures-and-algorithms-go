# LongestPeak in Go

This Go module provides multiple implementations for solving the **Longest Peak** problem, where a **peak** is defined as a contiguous sequence of strictly increasing integers followed by strictly decreasing integers, with a minimum length of 3.

## Problem Definition

Given an array of integers, a **peak** is a subarray where:

* There exists an element greater than both its neighbors (`array[i-1] < array[i] > array[i+1]`)
* The sequence before `i` is strictly increasing
* The sequence after `i` is strictly decreasing

The goal is to return the length of the **longest peak** found in the array.

---

## Functions

### `LongestPeak(array []int) int`

This is the primary and most efficient implementation. It traverses the array once, expanding left and right from each detected peak tip to find the full extent of the peak.

* **Time Complexity**: O(n)
* **Space Complexity**: O(1)

### `LongestPeak2(array []int) int`

An experimental and incomplete alternative. It detects all peak tips first and then attempts to expand from each, but contains logic errors:

* Fails to correctly calculate peak lengths.
* Does not update or track left and right bounds correctly.
* Returns incorrect results.

> **Note:** This version should not be used in production without corrections.

### `LongestPeak3(array []int) int`

This version is logically identical to `LongestPeak` but is written more cleanly and without external libraries like `math.Max`.

* **Time Complexity**: O(n)
* **Space Complexity**: O(1)
* Slightly more readable than version 1 for beginners.

---

## Example

```go
input := []int{1, 3, 2, 1, 2, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}
fmt.Println(LongestPeak(input))  // Output: 6
fmt.Println(LongestPeak3(input)) // Output: 6
```

---

## Running the Code

To run the code:

```bash
go run main.go
```

Make sure to provide the appropriate `main()` function if testing locally.

---

## License

This code is open-sourced under the MIT License.

---

Would you like me to help fix `LongestPeak2` to match the others?
