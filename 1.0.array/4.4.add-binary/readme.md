# Binary Addition in Go

## Overview
This Go function **`addBinary(a, b string) string`** takes two binary strings as input and returns their sum as a binary string. The function performs binary addition just like how we do manual binary addition, bit by bit, from right to left.

---

## Implementation
### **Function Code**
```go
package main

import (
    "fmt"
    "strconv"
)

func addBinary(a, b string) string {
    i, j, carry := len(a)-1, len(b)-1, 0
    result := ""

    for i >= 0 || j >= 0 || carry == 1 {
        if i >= 0 {
            carry += int(a[i]) - '0'  // Convert char to int and add to carry
            i--
        }
        if j >= 0 {
            carry += int(b[j]) - '0'  // Convert char to int and add to carry
            j--
        }
        
        result = strconv.Itoa(carry%2) + result  // Prepend the binary digit
        carry /= 2  // Update carry
    }
    return result
}

func main() {
    fmt.Println(addBinary("101", "11")) // Output: "1000"
}
```

---

## **How It Works**
### **Key Steps**
1. **Start from the rightmost bit** of both strings.
2. **Extract each bit and convert to integer** (`'0'` → `0`, `'1'` → `1`).
3. **Perform bitwise addition**, tracking a `carry`.
4. **Prepend the result bit-by-bit** to form the final binary sum.
5. **Continue until all bits are processed**, including any remaining carry.

### **Understanding `carry += int(a[i]) - '0'`**
- `a[i]` gives a character like `'1'` or `'0'`.
- `int(a[i])` converts it to an ASCII value (`'0'` = 48, `'1'` = 49).
- Subtracting `'0'` (`48`) gives the actual binary value:
  - `'0'` → `48 - 48 = 0`
  - `'1'` → `49 - 48 = 1`

### **Understanding `result = strconv.Itoa(carry%2) + result`**
- `carry % 2` extracts the **least significant bit** of the current sum.
- `strconv.Itoa(carry % 2)` converts it to a string.
- The `+ result` **prepends** it to the final result, ensuring correct order.
- `carry /= 2` shifts the carry leftward for the next iteration.

---

## **Example Walkthrough**
### Input: `a = "101", b = "11"`
### **Manual Binary Addition**
```
    101  (5 in decimal)
  + 011  (3 in decimal)
  -------
   1000  (8 in decimal)
```

### **Step-by-Step Execution**
| i | j | carry | carry % 2 | Updated result | Updated carry |
|---|---|-------|-----------|---------------|---------------|
| 2 | 1 | 2     | 0         | "0"           | 1             |
| 1 | 0 | 2     | 0         | "00"          | 1             |
| 0 | -1| 2     | 0         | "000"         | 1             |
| -1| -1| 1     | 1         | "1000"        | 0             |

### **Final Output**
✅ `result = "1000"`

---

## **Alternative Approach: Using a Slice (Efficient)**
String concatenation (`+ result`) is **inefficient** because strings are immutable. Instead, we can **use a slice** and reverse it at the end:

```go
func addBinary(a, b string) string {
    i, j, carry := len(a)-1, len(b)-1, 0
    var result []byte

    for i >= 0 || j >= 0 || carry > 0 {
        if i >= 0 {
            carry += int(a[i] - '0')
            i--
        }
        if j >= 0 {
            carry += int(b[j] - '0')
            j--
        }

        result = append(result, byte(carry%2)+'0')
        carry /= 2
    }

    // Reverse the result slice
    for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
        result[left], result[right] = result[right], result[left]
    }
    return string(result)
}
```

### **Why This is Better?**
- Avoids **expensive string concatenation** (`+ result` in a loop is slow).
- Uses a **byte slice**, which is more memory-efficient.
- Reverses the result **at the end**, making it optimal.

---

## **Time & Space Complexity**
### **Time Complexity: O(n)**
- We iterate through the **longest** input string **once**.
- Each operation (addition, modulo, division) is **O(1)**.

### **Space Complexity: O(n)**
- The result string grows **proportionally** to the longest input.
- The slice-based approach avoids unnecessary string copies, making it more memory-efficient.

---

## **Summary**
✅ **Works like manual binary addition** (right to left, bit-by-bit).  
✅ **Handles different string lengths and carries correctly.**  
✅ **Prepending result ensures correct bit order.**  
✅ **Optimized version using slices is more efficient.**  
✅ **Time Complexity O(n), Space Complexity O(n).**  

### **Example Run**
```go
fmt.Println(addBinary("101", "11")) // Output: "1000"
```
This function is a simple yet powerful way to handle binary addition efficiently in Go! 🚀


[source](https://leetcode.com/problems/add-binary/)