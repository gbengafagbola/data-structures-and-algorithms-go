package main

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Generate6DigitRandomInt generates a 6-digit random integer
func Generate6DigitRandomInt() int64 {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	return RandomInt(100000, 999999)
}

func main() {
	// Generate a 6-digit random integer
	randomInt := Generate6DigitRandomInt()
	fmt.Println(randomInt)
}
