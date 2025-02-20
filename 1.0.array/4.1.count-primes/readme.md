# Visual Explanation of `countPrimes(n)`

This function counts the number of prime numbers less than `n` using the **Sieve of Eratosthenes**. Let’s break it down visually.

---

## Algorithm Steps

1. Create a boolean array `isPrime` of size `n`, initialized to `true`.
2. Mark `0` and `1` as `false` since they are not prime.
3. Start from `2` and mark all multiples as `false`.
4. Continue for numbers up to `sqrt(n)`.
5. Count the remaining `true` values.

---

## Step-by-Step Execution

### **Step 1: Initialize an Array**
We create a boolean array `isPrime` of size `n` and set all values to `true`, assuming all numbers are prime.

```
Index:   0   1   2   3   4   5   6   7   8   9  10
Value:   F   F   T   T   T   T   T   T   T   T   T
```
*(0 and 1 are marked **false** since they are not prime.)*

### **Step 2: Start Sieving from 2**
For each prime `p`, mark all multiples of `p` as non-prime.

- **Mark multiples of 2**:
```
Index:   0   1   2   3   4   5   6   7   8   9  10
Value:   F   F   T   T   F   T   F   T   F   T   F
```
*(Numbers `4, 6, 8, 10` are marked as non-prime.)*

- **Mark multiples of 3**:
```
Index:   0   1   2   3   4   5   6   7   8   9  10
Value:   F   F   T   T   F   T   F   T   F   F   F
```
*(Number `9` is marked as non-prime.)*

### **Step 3: Count the Remaining `true` Values**
After sieving, the remaining `true` values represent prime numbers: **2, 3, 5, 7**.

So, `countPrimes(10) = 4`.

---

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

---

## Why Use `sqrt(n)`?
Instead of checking all numbers up to `n`, we only iterate up to `sqrt(n)`. This is based on the mathematical property that **if a number is composite, it must have a factor less than or equal to `sqrt(n)`**.

### **Example: Finding Factors of 36**
Consider `n = 36`. Its factors are:
```
1 × 36
2 × 18
3 × 12
4 × 9
6 × 6   ⬅️ sqrt(36) = 6
9 × 4
12 × 3
18 × 2
36 × 1
```
Notice how after `6 × 6`, the factors just repeat in reverse order. This means we only need to check numbers up to `6` (which is `sqrt(36)`) to determine if `36` is composite.

---

## Optimizing the Loop Condition

### **Loop Condition:**
```go
for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
```

### **Breaking It Down:**
1. `math.Sqrt(float64(n))` → Computes the square root of `n`.
2. `float64(n)` → Ensures `n` is treated as a floating-point number.
3. `int(...)` → Converts the square root to an integer.
4. `+1` → Ensures we include the integer part of `sqrt(n)`.

### **Example Calculation: (n = 50)**
```go
math.Sqrt(float64(50))    // 7.071
int(7.071) + 1            // 8
```
So, the loop runs from `2` to `7`, optimizing performance by avoiding redundant checks.

---

## Performance Analysis

### **Why This Works**
1. If `n` is composite, it must have a factor ≤ `sqrt(n)`.
2. Any larger factor must be paired with a smaller factor already checked.
3. This significantly reduces the number of iterations.

### **Impact on Performance**
- **Without `sqrt(n)`:** O(n) time complexity (too slow for large `n`).
- **With `sqrt(n)`:** O(n log log n) using Sieve of Eratosthenes (much faster).

By using `sqrt(n)`, we cut unnecessary iterations, making prime checking efficient! 🚀

---

## Complexity Analysis
- **Time Complexity:** O(n log log n) using the Sieve of Eratosthenes.
- **Space Complexity:** O(n) for the boolean array.

---

## Summary
This implementation efficiently counts primes using the **Sieve of Eratosthenes**, and the `sqrt(n)` trick optimizes performance by limiting unnecessary iterations. 🚀


[source](https://https://leetcode.com/problems/count-primes)