package main
import "strings"

func reverseWords(s string) string {
    // Split the string by whitespace and ignore empty elements
    words := strings.Fields(s) // this removes extra spaces

    result := []string{}

    // Reverse loop — your original logic
    for i := len(words) - 1; i >= 0; i-- {
        result = append(result, words[i])
    }

    return strings.Join(result, " ")
}
