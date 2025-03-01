package main

import "fmt"

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




func main(){
	array := []int{1, 1, 6, 1}
	sequence:= []int{1, 1, 1, 6}

	fmt.Println(IsValidSubsequence(array, sequence))
}