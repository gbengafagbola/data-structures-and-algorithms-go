package main

import (
	"fmt"
)

type Replacement struct {
	index  int
	Before string
	After  string
}

func replace(sourceCode string, replacements []Replacement) string {
	// Iterate over each replacement
	for _, r := range replacements {
		// Check if the substring to be replaced exists at the specified index
		if r.index+len(r.Before) <= len(sourceCode) && sourceCode[r.index:r.index+len(r.Before)] == r.Before {
			// Perform the replacement
			sourceCode = sourceCode[:r.index] + r.After + sourceCode[r.index+len(r.Before):]
		}
	}
	return sourceCode
}

func main() {
	sourceCode := "Hi There"
	r := []Replacement{
		{index: 0, Before: "Hi", After: "Hello"},
		{index: 3, Before: "There", After: "World"},
	}
	fmt.Println(replace(sourceCode, r)) // Output: "Hello World"
}