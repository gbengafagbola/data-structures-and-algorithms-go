package main

import (
	"fmt"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	gcdLen := gcd(len(str1), len(str2))
	return str1[:gcdLen]
}

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))     // "ABC"
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))     // "AB"
	fmt.Println(gcdOfStrings("LEET", "CODE"))       // ""
	fmt.Println(gcdOfStrings("AAAAAA", "AA"))       // "AA"
}



func gcdOfStringsNonMatheaticalApproach(str1 string, str2 string) string {
    p1 := 0
    p2 := len(str2) - 1
    unitLength := len(str2)

    for p2 < len(str1) {
        if str1[p1:p2+1] != str2 {
            return ""
        }
        p1 += unitLength
        p2 += unitLength
    }

    if p1 == len(str1) {
        return str2
    }
    return ""
}
