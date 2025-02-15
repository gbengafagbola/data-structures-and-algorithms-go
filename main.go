package main

import "fmt"

import (
	"dsa-golang/generate"
	"log" // Importing log package
)

func main() {
	// Call the Generate function and handle errors
	if err := generate.Generate(); err != nil {
		log.Fatalf("Error: %v", err)  // Log the error and exit
	}
}