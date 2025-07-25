package main
import "strings"


func reverseWords(s string) string {
    // Convert string to a byte slice to work efficiently
    n := len(s)
    var words []string
    i := 0

    for i < n {
        // Skip leading spaces
        for i < n && s[i] == ' ' {
            i++
        }
        if i >= n {
            break
        }

        // Extract word
        j := i
        for j < n && s[j] != ' ' {
            j++
        }
        words = append(words, s[i:j])
        i = j
    }

    // Reverse the words in-place
    for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
        words[i], words[j] = words[j], words[i]
    }

    return strings.Join(words, " ")
}



func reverseWordsInitial(s string) string {
    // Split the string by whitespace and ignore empty elements
    words := strings.Fields(s) // this removes extra spaces

    result := []string{}

    // Reverse loop — your original logic
    for i := len(words) - 1; i >= 0; i-- {
        result = append(result, words[i])
    }

    return strings.Join(result, " ")
}
