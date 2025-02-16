package main

import "fmt"

func isBadVersion(num int) bool {
	if num >= 3 {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	left := 1
	right := n
	for left < right {
		mid := left + (right-left)/2 //preventing overflow

		if isBadVersion(mid) {
			if mid == 1 || (mid-1 > 0 && !isBadVersion(mid-1)) {
				return mid
			} else {
				right = mid
			}
		}  else {
			left = mid + 1
		}

	}
	return left
}

func main() {
	num := 1
	fmt.Println(firstBadVersion(num))
}
