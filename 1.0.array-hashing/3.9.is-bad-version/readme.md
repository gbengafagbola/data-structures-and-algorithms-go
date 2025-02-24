Here’s a graphical representation of the **First Bad Version** problem using binary search.

---

### **Assumptions:**
- **Total Versions (`n`)** = 5  
- **First Bad Version** = 3  
- Function `isBadVersion(version)` returns `true` if the version is bad, otherwise `false`.

---

### **Step-by-Step Binary Search Visualization**

#### **Step 1: Start Binary Search**
- `left = 1`, `right = 5`
- Compute `mid = (1 + 5) / 2 = 3`
- `isBadVersion(3) = true` → Move `right` to `mid`

```
Versions:   1   2   3   4   5
Status:    [✅  ✅  ❌  ❌  ❌]  
             ↑       ↑
           left    right
           mid = 3 (Bad) → Move right to mid
```

---

#### **Step 2: Continue Search**
- `left = 1`, `right = 3`
- Compute `mid = (1 + 3) / 2 = 2`
- `isBadVersion(2) = false` → Move `left` to `mid + 1`

```
Versions:   1   2   3   4   5
Status:    [✅  ✅  ❌  ❌  ❌]  
                 ↑   ↑
               left  right
           mid = 2 (Good) → Move left to mid + 1
```

---

#### **Step 3: Converge to Answer**
- `left = 3`, `right = 3`  
- `left == right` → First bad version found at **3** 🎯

```
Versions:   1   2   3   4   5
Status:    [✅  ✅  ❌  ❌  ❌]  
                     ↑
                  left == right → First bad version found (3)
```

---

### **Final Output**
The first bad version is **3**. ✅

---

### **Binary Search Efficiency**
Instead of checking **each version one by one (`O(n)`)**, we use **binary search (`O(log n)`)**, making it significantly faster for large `n`.

---

### **Visual Summary**
```
Step 1: [✅  ✅  ❌  ❌  ❌]   mid = 3 (Bad) → Move right
Step 2: [✅  ✅  ❌  ❌  ❌]   mid = 2 (Good) → Move left
Step 3: [✅  ✅  ❌  ❌  ❌]   left == right → Found first bad version at 3
```


[source](https://leetcode.com/problems/first-bad-version)