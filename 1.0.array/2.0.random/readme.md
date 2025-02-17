# Random 6-Digit Integer Generator

## Overview
This Go program generates a random 6-digit integer using the built-in `math/rand` package. The number is printed to the console when the program runs.

## Features
- Generates a secure 6-digit random integer between `100000` and `999999`.
- Uses a time-based seed to ensure randomness.
- Implements modular functions for better readability and reusability.

## Code Explanation

```go
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
```

### Breakdown:
1. **`RandomInt(min, max int64) int64`**
   - Takes two integer arguments (`min` and `max`).
   - Uses `rand.Int63n()` to generate a random number in the given range.

2. **`Generate6DigitRandomInt()`**
   - Seeds the random number generator using `time.Now().UnixNano()` to ensure unique results.
   - Calls `RandomInt(100000, 999999)` to generate a 6-digit number.

3. **`main()` Function**
   - Calls `Generate6DigitRandomInt()` and prints the generated number to the console.

## Visual Representation
```
+--------------------------------+
|       Random Number Gen        |
+--------------------------------+
|                                |
|      [  123456  ]             |
|                                |
+--------------------------------+
```
The program outputs a randomly generated 6-digit number each time it is executed.

## How to Run
1. Install Go if not already installed: [Download Go](https://go.dev/dl/)
2. Save the code in a file named `main.go`.
3. Open a terminal and navigate to the file directory.
4. Run the command:
   ```sh
   go run main.go
   ```
5. The program prints a random 6-digit integer to the console.

## Use Cases
- OTP (One-Time Password) generation
- Unique identifier generation
- Randomized security codes

## Notes
- The `math/rand` package is not cryptographically secure. For secure applications, use the `crypto/rand` package instead.
