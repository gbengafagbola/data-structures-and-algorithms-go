# TwoNumberSum Function in Go

## Overview

The `TwoNumberSum` function identifies two distinct numbers within an integer array that sum up to a specified target value. This solution employs a hash map for efficient lookups, ensuring an optimal time complexity.

## Function Signature

```go
func TwoNumberSum(array []int, target int) []int
```

- **Parameters:**
  - `array`: A slice of integers to search within.
  - `target`: The integer sum to find within the array.

- **Returns:**
  - A slice containing the two integers that add up to the target. If no such pair exists, it returns an empty slice.

## Implementation Details

The function operates as follows:

1. **Initialization**: Create an empty hash map (`m`) to store numbers encountered during iteration.

2. **Iteration**: Traverse each number (`v`) in the array:
   - Compute the difference (`diff`) between the target and the current number (`diff = target - v`).
   - Check if `diff` exists in the hash map:
     - If it does, a pair has been found; return `[diff, v]`.
     - If not, add the current number (`v`) to the hash map for future reference.

3. **Completion**: If the loop concludes without finding a pair, return an empty slice.

## Example Usage

```go
package main

import (
    "fmt"
)

func TwoNumberSum(array []int, target int) []int {
    m := make(map[int]struct{})
    for _, v := range array {
        diff := target - v
        if _, ok := m[diff]; ok {
            return []int{diff, v}
        }
        m[v] = struct{}{}
    }
    return []int{}
}

func main() {
    array := []int{3, 5, -4, 8, 11, 1, -1, 6}
    target := 10
    result := TwoNumberSum(array, target)
    if len(result) == 0 {
        fmt.Println("No pair found.")
    } else {
        fmt.Printf("Pair found: %d + %d = %d\n", result[0], result[1], target)
    }
}
```

**Output:**
```
Pair found: 11 + -1 = 10
```

In this example, the function successfully identifies that `11` and `-1` sum up to the target value of `10`.

## Performance Analysis

- **Time Complexity**: O(n), where n is the number of elements in the array. Each element is processed once.
- **Space Complexity**: O(n), due to the storage requirements of the hash map.

This efficient approach ensures that even large datasets are handled promptly.

## Edge Cases

- **No Valid Pair**: If no two numbers sum to the target, the function returns an empty slice.
- **Multiple Pairs**: The function returns the first valid pair it encounters. If multiple pairs exist, subsequent pairs are not returned.

## Conclusion

The `TwoNumberSum` function offers an efficient solution to finding two numbers in an array that sum to a specified target. By leveraging a hash map for constant-time lookups, it ensures optimal performance suitable for large datasets.

For further reading and alternative approaches to this problem, consider exploring the following resources:

- [Two Sum Problem on LeetCode](https://leetcode.com/problems/two-sum/)
- [Solving LeetCode Two Sum: A Go Guide](https://medium.com/@AlexanderObregon/solving-the-two-sum-problem-on-leetcode-go-answer-s-walkthrough-d5a5f529a94c)

These articles provide in-depth discussions and varied solutions to the Two Sum problem, enhancing understanding and offering different perspectives. 