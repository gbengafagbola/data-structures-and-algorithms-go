package main

import "fmt"

// import (
// 	"dsa-golang/generate"
// 	"log" // Importing log package
// )

// func main() {
// 	// Call the Generate function and handle errors
// 	if err := generate.Generate(); err != nil {
// 		log.Fatalf("Error: %v", err)  // Log the error and exit
// 	}
// }

// sourceCode := "Hi There"
// r := replacement{index: 0, Before: "Hi", After: "Hello"}

type Replacement struct {
	index int
	Before string
	After string
}

func replace(sourceCode string, replacements []Replacement) string {

	// sourceCode = sourceCode[r.index:len(Before)]

	for _, r := range replacements {
		sourceCode = sourceCode[:r.index] + r.After + sourceCode[r.index+len(r.Before):]
	}


	return ""

}

func main(){
	sourceCode := "Hi There"
	r := []Replacement{{index: 0, Before: "Hi", After: "Hello"}, {index: 4, Before: "String", After: "After"}}
	fmt.Println(replace(sourceCode, r))
}

//  01234567
// "Hi There"