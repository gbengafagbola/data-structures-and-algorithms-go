package main

import "fmt"


func minWindow(s, t string) string {
	lenString := len(s)
	lenPattern := len(t)

	if lenString < lenPattern {
		return ""
	}

	hashString := [128]int{}
	hashPattern := [128]int{}

	for i := 0; i < lenPattern; i++ {
		hashPattern[t[i]]++
	}

	count, left, startIndex, minLen := 0, 0, -1, int(^uint(0)>>1) // or minLen := len(s) + 1  --> make it less complex for this case

	for right := 0; right < lenString; right++ {
		c := s[right]
		hashString[c]++

		if hashPattern[c] > 0 && hashString[c] <= hashPattern[c] {
			count++
		}

		for count == lenPattern {
			windowLength := right - left + 1
			if minLen > windowLength {
				minLen = windowLength
				startIndex = left
			}

			leftChar := s[left]
			hashString[leftChar]--

			if hashPattern[leftChar] > 0 && hashString[leftChar] < hashPattern[leftChar] {
				count--
			}
			left++
		}
	}

	if startIndex == -1 {
		return ""
	}
	return s[startIndex : startIndex+minLen]
}



func minWindow2(s, t string) string {

	lenString := len(s)
	lenPattern:= len(t)

	if lenString < lenPattern {
		return ""
	}

	hashString := make(map[byte]int)
	hashPattern := make(map[byte]int)

	for i := 0; i < lenPattern; i++ {
		hashPattern[t[i]]++
	}

	count := 0
	left := 0
	startIndex := -1
	minLen := int(^uint(0) >> 1) 

	for right := 0; right < lenString; right++ {
		c := s[right]
		hashString[c]++

		if patternCount, ok := hashPattern[c]; ok && hashString[c] <= patternCount {
			count++
		}

		if count == lenPattern{
			for hashPattern[s[left]] == 0 || hashString[s[left]] > hashPattern[s[left]] {
				leftCharacter := s[left]
				if hashString[leftCharacter] > 0 {
					hashString[leftCharacter]--
				}
				left++
			}
			windowLength := right - left +1
			if minLen > windowLength {
				minLen = windowLength
				startIndex = left
			}
		}
	}
	if 	startIndex == -1 {
		return ""
	}
	 return s[startIndex : startIndex + minLen]
}



func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t)) // Expected Output: "BANC"
}
