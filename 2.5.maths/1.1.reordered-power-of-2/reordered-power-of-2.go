package main

import (
	"fmt"
	"strconv"
)

func reorderedPowerOf2(N int) bool {
	// Convert N to slice of digits
	S := strconv.Itoa(N)
	A := make([]int, len(S))
	for i := 0; i < len(S); i++ {
		A[i] = int(S[i] - '0')
	}
	return permutations(A, 0)
}

// Check if slice of digits A represents a power of 2
func isPowerOfTwo(A []int) bool {
	if A[0] == 0 {
		return false // no leading zero
	}

	N := 0
	for _, x := range A {
		N = 10*N + x
	}

	// Remove factors of 2
	for N > 0 && (N&1) == 0 {
		N >>= 1
	}
	return N == 1
}

// Try all permutations of digits
func permutations(A []int, start int) bool {
	if start == len(A) {
		return isPowerOfTwo(A)
	}

	for i := start; i < len(A); i++ {
		swap(A, start, i)
		if permutations(A, start+1) {
			return true
		}
		swap(A, start, i) // backtrack
	}

	return false
}

func swap(A []int, i, j int) {
	A[i], A[j] = A[j], A[i]
}

func main() {
	fmt.Println(reorderedPowerOf2(128)) // true
	fmt.Println(reorderedPowerOf2(123)) // false
}
