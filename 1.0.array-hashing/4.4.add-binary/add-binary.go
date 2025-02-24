package main

import (
	"fmt"
	"strconv"
)

func addBinary(a, b string) string {

	i := len(a) -1
	j := len(b) -1

	carry := 0 
	result := ""

	for i >= 0 || j >= 0 || carry == 1 {

		if(i >= 0) {
			carry += int(a[i]) - '0'
			i--
		}

		if(j >= 0) {
			carry += int(b[j]) - '0'
			j--
		}

		result = strconv.Itoa(carry%2) + result
		carry /= 2
	}
	return result
}

func main() {
	a := "1111"
	b := "1101"
	fmt.Println(addBinary(a, b))
}