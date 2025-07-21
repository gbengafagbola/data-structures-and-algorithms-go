package main

import (
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	var builder strings.Builder
	builder.Grow(len(word1) + len(word2)) // 🔥 Preallocate memory

	i, j := 0, 0
	for i < len(word1) || j < len(word2) {
		if i < len(word1) {
			builder.WriteByte(word1[i])
			i++
		}
		if j < len(word2) {
			builder.WriteByte(word2[j])
			j++
		}
	}

	return builder.String()
}


func mergeAlternatelyNaive(word1 string, word2 string) string {
	i := 0
	var builder strings.Builder
	a := len(word1)
	b := len(word2)
	length := a
	check := word2

	if b < a {
		length = b
		check = word1
	}

	for i < length {
		builder.WriteString(string(word1[i]))
		builder.WriteString(string(word2[i]))
		i++
	}

	builder.WriteString(check[i:])
	return builder.String()
}
