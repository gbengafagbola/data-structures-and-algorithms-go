Here’s a concise but clear **README.md** for your `isPowerOfTwo` function:

---

# isPowerOfTwo – Go Implementation

## Overview

`isPowerOfTwo` is a Go function that determines whether a given integer is a power of two.
A number is considered a power of two if it can be expressed as `2^k` where `k` is a non-negative integer (e.g., `1, 2, 4, 8, 16, ...`).

---

## Function Signature

```go
func isPowerOfTwo(n int) bool
```

**Parameters**:

* `n` — The integer to check.

**Returns**:

* `true` if `n` is a power of two.
* `false` otherwise.

---

## How It Works

1. **Zero check** – If `n` is `0`, return `false` immediately since `0` is not a power of two.
2. **Division loop** – Repeatedly divide `n` by `2` as long as it is divisible by `2`.
3. **Final check** – If after division the result is exactly `1`, the original number was a power of two; otherwise, it wasn’t.

---

## Example Usage

```go
package main

import "fmt"

func isPowerOfTwo(n int) bool {
    if n == 0 {
        return false
    }
    for n%2 == 0 {
        n /= 2
    }
    return n == 1
}

func main() {
    fmt.Println(isPowerOfTwo(1))   // true  (2^0)
    fmt.Println(isPowerOfTwo(2))   // true  (2^1)
    fmt.Println(isPowerOfTwo(16))  // true  (2^4)
    fmt.Println(isPowerOfTwo(18))  // false
    fmt.Println(isPowerOfTwo(0))   // false
}
```

---

## Time & Space Complexity

* **Time Complexity:** O(log n) — because in the worst case, the function keeps dividing `n` by `2` until it becomes `1`.
* **Space Complexity:** O(1) — constant space usage.

---

## Alternative Approach

A faster bitwise approach avoids division:

```go
func isPowerOfTwoBitwise(n int) bool {
    return n > 0 && (n & (n - 1)) == 0
}
```

This works because powers of two have only **one** bit set in their binary representation.

