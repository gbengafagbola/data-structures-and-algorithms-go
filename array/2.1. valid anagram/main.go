package main

import "fmt"

func isAnagram(s string, t string) bool {

	// return false, if length of both string are not equal
	if len(s) != len(t) {
		return false
	}

	// create count of each unique character in string s
	// indexing with rune over astring allows proper handling og multi-byte characters useful in internationalized text
	// it's important to use rune to ensure that you're working with full characters, not just bytes hence rune over string.

	counts := make(map[rune]int)

	// loop through string s
	for _, char := range s {
		counts[char]++
	}

	// loop through string t and decrease the count for each character
	for _, char := range t {
		counts[char]--
		if counts[char] < 0 {
			// If a character count goes negative, the strings aren't anagrams
			return false
		}
	}

	// If all counts are zero, the strings are anagrams
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true

}

func isAnagramTwo(s string, t string) bool {
	// If the lengths of the strings are not equal, return false
	if len(s) != len(t) {
		return false
	}

	// Create a fixed-size array (26 elements for 'a' to 'z')
	//  characters are represented internally as numbers, based on their Unicode or ASCII values
	letters := make([]int, 26)
	// we can also use an array
	// letters := [26]int{0, 0, 0, ..., 0}

	// Loop through the first string and count occurrences of each character
	for i := 0; i < len(s); i++ {
		//The expression s[i] - 'a' calculates the position of the character in the alphabet, using 'a' as a reference point, where a = 0, b = 1, c = 2 ...

		// for example the ASCII of
		// The ASCII value of 'c' is 99.
		// The ASCII value of 'a' is 97.
		// So, 99 - 97 = 2.
		// Thus, s[i] - 'a' yields 2 for 'c'.

		letters[s[i]-'a']++ // Increment the count for each character in s
	}

	// Loop through the second string and decrement counts
	for i := 0; i < len(t); i++ {
		letters[t[i]-'a']-- // Decrement the count for each character in t
		if letters[t[i]-'a'] < 0 {
			// If any count goes negative, strings are not anagrams
			return false
		}
	}

	return true
}

func main() {
	anagram := isAnagram("mango", "goman")
	anagramTwo := isAnagramTwo("cat", "act")
	fmt.Println(anagram, anagramTwo)
}
