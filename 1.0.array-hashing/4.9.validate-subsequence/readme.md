# Subsequence Checker

This repository contains a Go implementation of a function that checks if a given `sequence` is a valid subsequence of an `array`. A subsequence is a sequence that can be derived from another sequence by deleting some or no elements without changing the order of the remaining elements.

## Code Explanation

### Function: `IsValidSubsequence`

```go
func IsValidSubsequence(array []int, sequence []int) bool {
	arrIndex := 0
	seqIndex := 0

	// Traverse both arrays
	for arrIndex < len(array) && seqIndex < len(sequence) {
		if array[arrIndex] == sequence[seqIndex] {
			seqIndex++ // Move to the next element in the sequence
		}
		arrIndex++ // Always move to the next element in the array
	}

	// If we've traversed the entire sequence, it's a valid subsequence
	return seqIndex == len(sequence)
}
```

### How It Works

1. **Two Pointers**:
   - `arrIndex`: Tracks the current position in the `array`.
   - `seqIndex`: Tracks the current position in the `sequence`.

2. **Traversal**:
   - The function iterates through both the `array` and the `sequence` simultaneously.
   - If the current element in the `array` matches the current element in the `sequence`, the `seqIndex` pointer is moved forward.
   - The `arrIndex` pointer is always moved forward, regardless of whether there is a match.

3. **Completion Check**:
   - If the `seqIndex` pointer reaches the end of the `sequence`, it means all elements of the `sequence` were found in the `array` in the correct order, and the function returns `true`.
   - Otherwise, it returns `false`.

### Example Walkthrough

#### Input:
```go
array := []int{5, 1, 22, 25, 6, -1, 8, 10}
sequence := []int{1, 6, -1, 10}
```

#### Visualization:

| Step | `arrIndex` | `seqIndex` | `array[arrIndex]` | `sequence[seqIndex]` | Action                     |
|------|------------|------------|-------------------|-----------------------|----------------------------|
| 1    | 0          | 0          | 5                 | 1                     | No match, move `arrIndex`  |
| 2    | 1          | 0          | 1                 | 1                     | Match, move both pointers  |
| 3    | 2          | 1          | 22                | 6                     | No match, move `arrIndex`  |
| 4    | 3          | 1          | 25                | 6                     | No match, move `arrIndex`  |
| 5    | 4          | 1          | 6                 | 6                     | Match, move both pointers  |
| 6    | 5          | 2          | -1                | -1                    | Match, move both pointers  |
| 7    | 6          | 3          | 8                 | 10                    | No match, move `arrIndex`  |
| 8    | 7          | 3          | 10                | 10                    | Match, move both pointers  |

#### Result:
- `seqIndex` reaches the end of the `sequence` (i.e., `seqIndex == len(sequence)`), so the function returns `true`.

### Another Example

#### Input:
```go
array := []int{5, 1, 22, 25, 6, -1, 8, 10}
sequence := []int{5, 1, 25, 22, 6, -1, 8, 10}
```

#### Visualization:

| Step | `arrIndex` | `seqIndex` | `array[arrIndex]` | `sequence[seqIndex]` | Action                     |
|------|------------|------------|-------------------|-----------------------|----------------------------|
| 1    | 0          | 0          | 5                 | 5                     | Match, move both pointers  |
| 2    | 1          | 1          | 1                 | 1                     | Match, move both pointers  |
| 3    | 2          | 2          | 22                | 25                    | No match, move `arrIndex`  |
| 4    | 3          | 2          | 25                | 25                    | Match, move both pointers  |
| 5    | 4          | 3          | 6                 | 22                    | No match, move `arrIndex`  |
| 6    | 5          | 3          | -1                | 22                    | No match, move `arrIndex`  |
| 7    | 6          | 3          | 8                 | 22                    | No match, move `arrIndex`  |
| 8    | 7          | 3          | 10                | 22                    | No match, move `arrIndex`  |

#### Result:
- `seqIndex` does not reach the end of the `sequence` (i.e., `seqIndex < len(sequence)`), so the function returns `false`.

### Usage

To use the `IsValidSubsequence` function, simply call it with your `array` and `sequence` as arguments:

```go
package main

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}

	if IsValidSubsequence(array, sequence) {
		println("Valid subsequence")
	} else {
		println("Not a valid subsequence")
	}
}
```

### Output:
```
Valid subsequence
```

## Conclusion

This implementation ensures that the `sequence` is a valid subsequence of the `array` by maintaining the order of elements. The two-pointer approach is efficient and works in **O(n) time complexity**, where `n` is the length of the `array`.