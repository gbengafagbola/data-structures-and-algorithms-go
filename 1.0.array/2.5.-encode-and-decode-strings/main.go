package main

import (
	"encoding/json" 
	"fmt"
	"strconv"
	"strings"
)

// Encode a list of strings to a single string
func encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		// Prefix the string with its length and a colon
		sb.WriteString(fmt.Sprintf("%d:%s", len(str), str))
	}
	return sb.String()
}

// Decode a single string back to a list of strings
func decode(encoded string) []string {
	result := []string{}
	i := 0
	for i < len(encoded) {
		// Find the colon to separate length and string
		colonIdx := strings.Index(encoded[i:], ":")
		if colonIdx == -1 {
			break // No more encoded strings
		}

		// Extract the length of the next string
		length, err := strconv.Atoi(encoded[i : i+colonIdx])
		if err != nil {
			panic("Invalid encoding format")
		}

		// Extract the string using the parsed length
		start := i + colonIdx + 1
		end := start + length
		result = append(result, encoded[start:end])

		// Move to the next encoded string
		i = end
	}
	return result
}

// Utility to format a slice with each string wrapped in double quotes
func formatSliceWithQuotes(slice []string) string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, str := range slice {
		sb.WriteString(fmt.Sprintf("\"%s\"", str))
		if i < len(slice)-1 {
			sb.WriteString(", ") // Add a comma and space except for the last element
		}
	}
	sb.WriteString("]")
	return sb.String()
}


// Encode a list of strings to a single string
func encodeJSON(strs []string) string {
	encoded, _ := json.Marshal(strs)
	return string(encoded)
}

// Decode a single string back to a list of strings
func decodeJSON(encoded string) []string {
	var result []string
	json.Unmarshal([]byte(encoded), &result)
	return result
}

func main() {
	// Example Usage
	strs := []string{"neet", "code", "love", "you"}
	encoded := encode(strs)
	fmt.Println("Encoded:", encoded)

	decoded := decode(encoded)
	fmt.Println("Decoded:", formatSliceWithQuotes(decoded))
}
