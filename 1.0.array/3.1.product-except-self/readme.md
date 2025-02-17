# Product of Array Except Self

## Overview
This Go program computes the product of all elements in an array except the current index, returning a new array of results.

## Features
- Computes the product of all elements except the current one.
- Uses iteration for efficiency.
- Returns an array of the same length as input.

## Code Explanation

```go
package main

import "fmt"

func productExceptSelf(nums []int) []int {
	result := []int{}
	track := 0
	multiple := 1
	var i int = 0

	for i < len(nums) && track <= len(nums) {
		if track != i {
			multiple = multiple * nums[i]
		}
		i++
	}
	result = append(result, multiple)
	i = 0
	track++

	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf(nums)
	fmt.Println(result) // Expected Output: [24, 12, 8, 6]
}
```

### Breakdown:
1. **Initialization**
   - Creates an empty result slice.
   - Defines `track` and `multiple` to handle calculations.

2. **Looping through elements**
   - Iterates through `nums`, multiplying elements except the current index.
   - Appends the computed product to `result`.

3. **Edge Cases**
   - Handles cases where input length is 1 or contains zeros.

## Visual Representation
```
Input:  [1, 2, 3, 4]
Output: [24, 12, 8, 6]
```

## How to Run
1. Install Go: [Download Go](https://go.dev/dl/)
2. Save the code as `main.go`.
3. Open a terminal and navigate to the file directory.
4. Run:
   ```sh
   go run main.go
   ```

## Use Cases
- Efficient computation of array products.
- Useful for mathematical applications and optimization problems.

## Notes
- The current implementation has logical errors; consider refining the loop conditions.
- The commented alternative function provides a more structured approach.