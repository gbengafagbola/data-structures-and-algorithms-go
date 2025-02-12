package main

import (
	"fmt"
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	// Create a map to track seen letters
	seen := make(map[rune]bool)

	// Convert the word to lowercase for case-insensitivity
	word = strings.ToLower(word)

	for _, char := range word {
		// Skip non-alphabetic characters (spaces, hyphens, etc.)
		if !unicode.IsLetter(char) {
			continue
		}

		// Check if the character has been seen before
		if seen[char] {
			return false
		}

		// Mark the character as seen
		seen[char] = true
	}

	return true
}

func main() {
	// Test cases
	words := []string{"lumberjacks", "background", "downstream", "six-year-old", "isograms"}
	for _, word := range words {
		if IsIsogram(word) {
			fmt.Printf("%q is an isogram.\n", word)
		} else {
			fmt.Printf("%q is not an isogram.\n", word)
		}
	}
}
