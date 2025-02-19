package main

import (
	"fmt"
	"math"
)

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}

	isPrime := make([]bool, n)

	for i := 0; i < n; i++ {
		isPrime[i] = true
	}

	isPrime[0] = false
	isPrime[1] = false

	for i := 2; i < int(math.Ceil(math.Sqrt(float64(n)))); i++ {
		if isPrime[i] {
			for multiple_of_i := i * 2; multiple_of_i < n; multiple_of_i += i {
				isPrime[multiple_of_i] = false
			}
		}
	}

	primeCount := 0
	for i := 0; i < n; i++ {
		if isPrime[i] {
			primeCount++
		}
	}

	return primeCount
}


func main() {
	fmt.Println(countPrimes(20)) // Output: 8 (Primes: 2, 3, 5, 7, 11, 13, 17, 19)
}