package main

import (
	"log"       // Importing log package
	"dsa-golang/generate"
)

func main() {
	// Call the Generate function and handle errors
	if err := generate.Generate(); err != nil {
		log.Fatalf("Error: %v", err)  // Log the error and exit
	}
}
