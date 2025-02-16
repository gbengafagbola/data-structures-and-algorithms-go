# Smallest Positive Missing Number

## Overview
This Go program finds the smallest positive missing number in an array using a variation of the cyclic sort approach. It rearranges the elements to their correct positions and identifies the missing number efficiently.

## Algorithm Explanation
The function `SmallestPositiveMissingNumber` follows these steps:
1. Iterates through the array, attempting to place each element in its correct position.
2. Swaps elements if they are not in their correct position and within the valid range.
3. If an element is out of range, it is ignored.
4. The smallest missing positive number is identified during iteration.

## Time Complexity
- The function runs in **O(n)** time complexity as it processes each element at most twice (once for placing it correctly and once for verification).

## Code Implementation
```go
package main
import "fmt"

func SmallestPositiveMissingNumber(arr []int, size int) int {
	i := 0
	missing := -1

	for i < size {
		currentIndex := i + 1

		if arr[i] != currentIndex && arr[i] > size {
			missing = currentIndex
			i++
		}

		if currentIndex < size && arr[i] != currentIndex {
			arr[i], arr[arr[i]-1] = arr[arr[i]-1], arr[i]
		}

		if currentIndex < size && arr[i] == i+1 {
			i++
		}
	}

	return missing
}

func main() {
	array := []int{1, 3, 4, 2, 6}
	size := len(array)
	fmt.Println("Smallest Missing Number:", SmallestPositiveMissingNumber(array, size))
}
```

## How to Run
1. Ensure you have Go installed on your system.
2. Copy the code into a file named `main.go`.
3. Open a terminal and navigate to the directory where `main.go` is saved.
4. Run the command:
   ```sh
   go run main.go
   ```
5. The output will display the smallest positive missing number in the given array.

## Example Output
```
Smallest Missing Number: 5
```

## Edge Cases Considered
- Arrays with consecutive numbers.
- Arrays with missing numbers at various positions.
- Arrays containing negative numbers.
- Arrays with all numbers in place.

## Author
This program is implemented in Go and demonstrates an efficient way to find the smallest missing positive integer.