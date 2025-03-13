package main

import (
	"fmt"
	"strings"
)

type replacement struct {
	index  int
	Before string
	After  string
}

// Method to perform replacement
func (r replacement) replace(sourceCode string) string {
	if strings.HasPrefix(sourceCode[r.index:], r.Before) {
		return sourceCode[:r.index] + r.After + sourceCode[r.index+len(r.Before):]
	}
	return sourceCode
}

func main() {
	sourceCode := "Hi There"
	r := replacement{index: 0, Before: "Hi", After: "Hello"}
	newCode := r.replace(sourceCode)
	fmt.Println(newCode) // Output: "Hello There"
}
