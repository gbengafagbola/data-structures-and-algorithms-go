package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}

func main() {
	// Test cases
	years := []int{1996, 1997, 1900, 2000}
	for _, year := range years {
		if isLeapYear(year) {
			fmt.Printf("%d is a leap year.\n", year)
		} else {
			fmt.Printf("%d is not a leap year.\n", year)
		}
	}
}