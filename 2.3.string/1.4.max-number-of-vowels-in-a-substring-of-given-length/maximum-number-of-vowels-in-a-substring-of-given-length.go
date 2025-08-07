package main

var vowels = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

func maxVowels(s string, k int) int {
	count := 0

	// Initial window
	for i := 0; i < k; i++ {
		if vowels[s[i]] {
			count++
		}
	}

	maxCount := count

	// Sliding window
	for i := k; i < len(s); i++ {
		if vowels[s[i]] {
			count++
		}
		if vowels[s[i-k]] {
			count--
		}
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
