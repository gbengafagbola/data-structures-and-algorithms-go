package main


func numOfUnplacedFruits(fruits []int, baskets []int) int {
	count := 0
	n := len(baskets)
	for _, fruit := range fruits {
		unset := 1
		for i := 0; i < n; i++ {
			if fruit <= baskets[i] {
				baskets[i] = 0
				unset = 0
				break
			}
		}
		count += unset
	}
	return count
}