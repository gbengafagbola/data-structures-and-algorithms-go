// FormatPhoneNo(phone, format)
// input = +2347037076000, 2347037076000, 07037076000
// formart '+234', '234', '0'
// output = +2347037076000, 2347037076000, 07037076000 or err

// FormatPhoneNo(phone, format) (formattedPhoneno, error)

package main

import (
	"errors"
	"fmt" 
)

func FormatPhoneNo(phone string, format string) (string, error) {
	// input = +2347037076000
	if len(phone) < 11 {
		return "", errors.New("invalid Number")
	}
	newPhoneNumber := ""
	phoneNumber := ""

		if phone[0] == '+' {
			phoneNumber = phone[4:]
		} else if (phone[0] == '2') {
			phoneNumber = phone[3:]
		} else if (phone[0] == '0') {
			phoneNumber = phone[1:]
		}

		newPhoneNumber = format + phoneNumber[1:]
	
		return newPhoneNumber, nil
}


func main() {
	phone := "+2347037076000"
	format := "234"
	// phoneNumber, err := FormatPhoneNo(phone, format)
	fmt.Println(FormatPhoneNo(phone, format))
}