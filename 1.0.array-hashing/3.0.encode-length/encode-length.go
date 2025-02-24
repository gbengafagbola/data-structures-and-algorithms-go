package main

import (
	"fmt" 
	"strconv"
	"strings" 
)

// RunLengthEncode performs run-length encoding on the input string.
func RunLengthEncode(input string) string {
	if len(input) == 0 {
		return ""
	}

	var result strings.Builder
	count := 1

	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			if count > 1 {
				result.WriteString(strconv.Itoa(count))
			}
			result.WriteByte(input[i-1])
			count = 1
		}
	}

	// Handle the last run
	if count > 1 {
		result.WriteString(strconv.Itoa(count))
	}
	result.WriteByte(input[len(input)-1])

	return result.String()
}

// RunLengthDecode performs run-length decoding on the input string.
func RunLengthDecode(input string) string {
	var result strings.Builder
	count := 0

	for i := 0; i < len(input); i++ {
		if input[i] >= '0' && input[i] <= '9' {
			n, _ := strconv.Atoi(string(input[i]))
			count = count*10 + n
		} else {
			if count == 0 {
				count = 1
			}
			result.WriteString(strings.Repeat(string(input[i]), count))
			count = 0
		}
	}

	return result.String()
}

func main() {
	// Example usage
	original := "WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB"
	encoded := RunLengthEncode(original)
	decoded := RunLengthDecode(encoded)

	fmt.Printf("Original: %s\n", original)
	fmt.Printf("Encoded: %s\n", encoded)
	fmt.Printf("Decoded: %s\n", decoded)

	// Another example
	original = "AABCCCDEEEE"
	encoded = RunLengthEncode(original)
	decoded = RunLengthDecode(encoded)

	fmt.Printf("Original: %s\n", original)
	fmt.Printf("Encoded: %s\n", encoded)
	fmt.Printf("Decoded: %s\n", decoded)
}
