package main

type Solution struct {

}


func (s *Solution) solution(digits string, digitToString map[rune]string, cur string, ans *[]string, digitToIndex int) {

	if len(cur) == len(digits) {
		*ans = append(*ans, cur)
		return
	}
	
	currentDigit := rune(digits[digitToIndex])
	currentString := digitToString[currentDigit]

	for _, char := range currentString {
		s.solution(digits, digitToString, cur+string(char), ans, digitToIndex+1)
	}

}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	digitToString := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	ans := []string{}
	s := Solution{}
	s.solution(digits, digitToString, "", &ans, 0)
	return ans
}