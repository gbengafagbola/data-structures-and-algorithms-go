### **Visual Explanation of `countPrimes(n)`**

This function counts the number of prime numbers less than `n` using the **Sieve of Eratosthenes**. Let’s break it down visually.

#### **Step 1: Initialize an array**
We create a boolean array `isPrime` of size `n` and set all values to `true`, assuming all numbers are prime.

```
Index:   0   1   2   3   4   5   6   7   8   9  10
Value:   F   F   T   T   T   T   T   T   T   T   T
```
(0 and 1 are marked **false** since they are not prime.)

#### **Step 2: Start sieving from 2**
For each prime `p`, mark all multiples of `p` as non-prime.

- **Mark multiples of 2**:
```
Index:   0   1   2   3   4   5   6   7   8   9  10
Value:   F   F   T   T   F   T   F   T   F   T   F
```
(Numbers `4, 6, 8, 10` are marked as non-prime.)

- **Mark multiples of 3**:
```
Index:   0   1   2   3   4   5   6   7   8   9  10
Value:   F   F   T   T   F   T   F   T   F   F   F
```
(Numbers `9` is marked as non-prime.)

#### **Step 3: Count the remaining `true` values**
After sieving, the remaining `true` values represent prime numbers: **2, 3, 5, 7**.  
So, `countPrimes(10) = 4`.

---

### **README Example**
Here's a README file with another example:  

```markdown
# Count Primes - Sieve of Eratosthenes

## Overview
This Go program counts the number of prime numbers less than `n` using the Sieve of Eratosthenes algorithm.

## Algorithm Steps
1. Create a boolean array `isPrime` of size `n` initialized to `true`.
2. Mark `0` and `1` as `false` since they are not prime.
3. Start from `2` and mark all multiples as `false`.
4. Continue for numbers up to `sqrt(n)`.
5. Count the remaining `true` values.

## Example Usage

```go
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

	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if isPrime[i] {
			for multiple_of_i := i * i; multiple_of_i < n; multiple_of_i += i {
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
```

## Time Complexity
- **O(n log log n)** using the Sieve of Eratosthenes.

## Space Complexity
- **O(n)** for the boolean array.

```

This README explains the function with an example and includes a corrected version of your function.  
Let me know if you need a different format! 🚀