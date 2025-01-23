package main

import "fmt"

func validAnagrams(s,t string) bool {
	a := make(map[rune]int)

	for _, i := range s {
		 a[i]++
	}

	for _,i := range t {
		a[i]++
   }


   for _, v := range a {
		if (v % 2 != 0) {
			return false
		}
   }


	return true
}

func main() {
	s := "racecar" 
	t := "carrace"
	fmt.Println("the array returns:", validAnagrams(s, t))
}
