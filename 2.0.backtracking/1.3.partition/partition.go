package main

func isPalindrome(s string) bool {
	l := 0
	r := len(s) -1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r++
	}
	return true

}


func solution(s string, curArr []string, ans *[][]string) {
	if len(s) == 0 {
		temp := make([]string, len(curArr))
		copy(temp, curArr)
		(*ans) = append((*ans), temp)
		return
	}

	for i := 1; i <= len(s); i++ {
		curString := s[0:i]
		if isPalindrome(curString) {
			curArr = append(curArr, curString)
			solution(s[i:], curArr, ans)
			curArr = curArr[:len(curArr)-1]
		}
	}
}

func partition(s string) [][]string {
	ans := [][]string{}
	solution(s, []string{}, &ans)
	return ans
}
