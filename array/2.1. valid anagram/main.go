package main

import "fmt"

func isAnagram(s string, t string) bool {

	// return false, if length of both string are not equal
	if len(s) != len(t) {
		return false
	}

	counts := make(map[int]int)

	for i := range s {
		counts[i]++
		if counts[i] > 1 {
			return true
		}

		
	}

	return false

}

func main() {
	anagram := isAnagram("mango", "goman")
	fmt.Println(anagram)
}
