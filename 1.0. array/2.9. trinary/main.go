package main

import (
	"fmt"
	"math" 
)

// ConvertTrinaryToDecimal converts a trinary string to its decimal equivalent.
// If the input is invalid, it returns 0.
func ConvertTrinaryToDecimal(trinary string) int {
	decimal := 0
	for i, char := range trinary {
		// Validate that the character is a trinary digit (0, 1, or 2).
		if char < '0' || char > '2' {
			return 0
		}
		// Convert character to integer.
		digit := int(char - '0')
		// Calculate the power of 3 for the current position.
		power := len(trinary) - 1 - i
		// Add the value of the current digit to the decimal result.
		decimal += digit * int(math.Pow(3, float64(power)))
	}
	return decimal
}

func main() {
	// Example usage
	trinary := "102012"
	result := ConvertTrinaryToDecimal(trinary)
	fmt.Printf("The decimal equivalent of trinary %s is %d\n", trinary, result)

	// Invalid trinary example
	invalidTrinary := "12345"
	result = ConvertTrinaryToDecimal(invalidTrinary)
	fmt.Printf("The decimal equivalent of invalid trinary %s is %d\n", invalidTrinary, result)
}
