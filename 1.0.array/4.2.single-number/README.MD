# 🔥 Finding the Unique Number Using XOR

## 🚀 Problem Statement
Given an array of integers where every element appears **twice except for one**, find the unique number in **O(n) time and O(1) space**.

### ✅ XOR Properties That Make This Work
1. **Identity Property:**  
   \[ x \oplus 0 = x \]  
   (XOR with 0 keeps the number unchanged.)

2. **Self-Canceling Property:**  
   \[ x \oplus x = 0 \]  
   (XOR of a number with itself is always 0.)

3. **Commutative & Associative:**  
   \[ a \oplus b = b \oplus a \]  
   \[ (a \oplus b) \oplus c = a \oplus (b \oplus c) \]  
   (Order of operations doesn’t matter.)

These properties allow all **duplicate numbers to cancel out**, leaving only the unique number.

---

## 💻 Code Implementation in Go

```go
func singleNumber(nums []int) int {
    answer := 0
    for _, num := range nums {
        answer ^= num
    }
    return answer
}
```

---

## 🔎 Step-by-Step Example

### **Example 1: nums = [4, 1, 2, 1, 2]**

| Step | XOR Operation | Result (`answer`) |
|------|-------------|-----------------|
| Init | `answer = 0` | 0 |
| `4`  | `0 ⊕ 4`    | **4** |
| `1`  | `4 ⊕ 1`    | **5** |
| `2`  | `5 ⊕ 2`    | **7** |
| `1`  | `7 ⊕ 1`    | **6** |
| `2`  | `6 ⊕ 2`    | **4** |
| **Final Answer** | **4** (Unique number) |

All duplicates (`1,1` and `2,2`) cancel out, leaving **4** as the only number.

---

### **Example 2: nums = [3, 3, 7, 2, 2]**

| Step | XOR Operation | Result (`answer`) |
|------|-------------|-----------------|
| Init | `answer = 0` | 0 |
| `3`  | `0 ⊕ 3`    | **3** |
| `3`  | `3 ⊕ 3`    | **0** |
| `7`  | `0 ⊕ 7`    | **7** |
| `2`  | `7 ⊕ 2`    | **5** |
| `2`  | `5 ⊕ 2`    | **7** |
| **Final Answer** | **7** (Unique number) |

Again, the duplicate numbers cancel out, leaving **7**.

---

## 🛠️ How the XOR Results Are Computed  
The XOR (`^`) operation works at the **bitwise level**, comparing corresponding bits of two numbers and returning `1` if the bits are different and `0` if they are the same.

### **Example 1: `nums = [4, 1, 2, 1, 2]`**

| Step | Number | Binary Representation | XOR Operation | Result |
|------|--------|----------------------|--------------|--------|
| Init | `answer = 0` | `0000` | - | `0000` |
| `4`  | `4`  | `0100` | `0000 ⊕ 0100` | `0100` (4) |
| `1`  | `1`  | `0001` | `0100 ⊕ 0001` | `0101` (5) |
| `2`  | `2`  | `0010` | `0101 ⊕ 0010` | `0111` (7) |
| `1`  | `1`  | `0001` | `0111 ⊕ 0001` | `0110` (6) |
| `2`  | `2`  | `0010` | `0110 ⊕ 0010` | `0100` (4) |

Final answer = **4** ✅ (since all duplicates canceled out)

---

### **Example 2: `nums = [3, 3, 7, 2, 2]`**

| Step | Number | Binary Representation | XOR Operation | Result |
|------|--------|----------------------|--------------|--------|
| Init | `answer = 0` | `0000` | - | `0000` |
| `3`  | `3`  | `0011` | `0000 ⊕ 0011` | `0011` (3) |
| `3`  | `3`  | `0011` | `0011 ⊕ 0011` | `0000` (0) |
| `7`  | `7`  | `0111` | `0000 ⊕ 0111` | `0111` (7) |
| `2`  | `2`  | `0010` | `0111 ⊕ 0010` | `0101` (5) |
| `2`  | `2`  | `0010` | `0101 ⊕ 0010` | `0111` (7) |

Final answer = **7** ✅

---

## 🎯 Why This is Efficient
- **Time Complexity:** `O(n)`, since we iterate through the list once.
- **Space Complexity:** `O(1)`, as we use only a single variable (`answer`).

### 🔥 Summary of How XOR Works  
1. **Same numbers cancel out (`x ⊕ x = 0`)**, removing duplicates.  
2. **Order doesn’t matter** due to commutative property.  
3. **The only number left is the unique one.**  

This is why XOR is a **powerful trick** for problems involving pairs! 🚀

