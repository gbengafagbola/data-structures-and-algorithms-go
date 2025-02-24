# Minimum Window Substring - Simplified Explanation

## Problem Statement
Given two strings **s** (a big string) and **t** (a small string), find the **smallest substring** in **s** that contains all characters of **t** in any order.

### Example
```
s = "ADOBECODEBANC"
t = "ABC"
Output: "BANC"
```
---

## **Step-by-Step Solution (Visual Explanation)**

### **1️⃣ Two Pointers (Left & Right)**
We use **two pointers** (`left` and `right`) to **move a window** across the string `s`.

- Start `right` at **0** and move it forward until all characters in `t` are in the window.
- Move `left` forward to **shrink** the window as much as possible while still keeping all `t` characters.
- Keep track of the **smallest valid window** found.

### **2️⃣ Hash Tables to Count Letters**
We use two **hash tables (arrays/maps)**:
- `hashPattern`: Counts the occurrences of each letter in `t`.
- `hashString`: Tracks the count of letters in the current window of `s`.

### **3️⃣ Move Right to Expand Until We Cover `t`**
We move `right` forward **until all characters of `t` are found** in `s`.

#### Example Progression:
```
s = "ADOBECODEBANC"
t = "ABC"

Step 1: Move `right` until all "ABC" are included
[ADOBEC]
     ^    ^
   left  right
```

### **4️⃣ Move Left to Find the Smallest Valid Window**
Once we have all `t` characters, we **move `left`** to make the window **as small as possible**.

#### Example Shrinking:
```
s = "ADOBECODEBANC"
t = "ABC"

Valid windows:
[ADOBECODEBANC] -> too large
     [BANC] -> smallest valid window ✅
```

---

## **Implementation 1: Using Fixed-Size Arrays** (Efficient Solution)
```go
package main

import "fmt"

func minWindow(s, t string) string {
    lenString, lenPattern := len(s), len(t)
    if lenString < lenPattern {
        return ""
    }

    hashString := [128]int{}
    hashPattern := [128]int{}

    for i := 0; i < lenPattern; i++ {
        hashPattern[t[i]]++
    }

    count, left, startIndex, minLen := 0, 0, -1, int(^uint(0)>>1)

    for right := 0; right < lenString; right++ {
        c := s[right]
        hashString[c]++

        if hashPattern[c] > 0 && hashString[c] <= hashPattern[c] {
            count++
        }

        for count == lenPattern {
            windowLength := right - left + 1
            if minLen > windowLength {
                minLen = windowLength
                startIndex = left
            }
            leftChar := s[left]
            hashString[leftChar]--
            if hashPattern[leftChar] > 0 && hashString[leftChar] < hashPattern[leftChar] {
                count--
            }
            left++
        }
    }

    if startIndex == -1 {
        return ""
    }
    return s[startIndex : startIndex+minLen]
}

func main() {
    fmt.Println(minWindow("ADOBECODEBANC", "ABC")) // Expected Output: "BANC"
}
```

### **Why This Works Efficiently?**
✅ **Fixed-Size Array (O(1) space)** instead of maps.
✅ **Each character is processed at most twice (O(n) time complexity)**.

---

## **Implementation 2: Using Hash Maps (Alternative Approach)**
```go
package main

func minWindow2(s, t string) string {
    lenString, lenPattern := len(s), len(t)
    if lenString < lenPattern {
        return ""
    }

    hashString := make(map[byte]int)
    hashPattern := make(map[byte]int)

    for i := 0; i < lenPattern; i++ {
        hashPattern[t[i]]++
    }

    count, left, startIndex, minLen := 0, 0, -1, int(^uint(0)>>1)

    for right := 0; right < lenString; right++ {
        c := s[right]
        hashString[c]++

        if patternCount, ok := hashPattern[c]; ok && hashString[c] <= patternCount {
            count++
        }

        if count == lenPattern {
            for hashPattern[s[left]] == 0 || hashString[s[left]] > hashPattern[s[left]] {
                leftCharacter := s[left]
                if hashString[leftCharacter] > 0 {
                    hashString[leftCharacter]--
                }
                left++
            }
            windowLength := right - left + 1
            if minLen > windowLength {
                minLen = windowLength
                startIndex = left
            }
        }
    }
    if startIndex == -1 {
        return ""
    }
    return s[startIndex : startIndex+minLen]
}
```

### **Differences Between Implementations**
| Feature              | Fixed-Size Array Implementation | Hash Map Implementation |
|----------------------|--------------------------------|-------------------------|
| Memory Usage        | O(1) (128 elements)            | O(k) (depends on unique chars) |
| Speed               | Faster (array lookup)          | Slightly slower (map lookup) |
| Simplicity          | More optimized                 | More flexible |

---

## **Final Recap**
- ✅ Use **Two Pointers (Sliding Window)** to scan `s` and find the **smallest** valid window.
- ✅ Use **Hash Tables (Arrays/Maps)** to track character counts.
- ✅ Move **Right Pointer** until all `t` characters are found, then move **Left Pointer** to shrink the window.
- ✅ **Return the smallest valid window**.

---

## **🎯 Best Choice?**
If performance matters ⏩ **Use Fixed-Size Arrays (Faster & Less Memory)**.
If `t` has a **large character set** (like Unicode) ⏩ **Use Hash Maps**.

🚀 **Happy Coding!** 🎉

