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

## 🎯 Why This is Efficient
- **Time Complexity:** `O(n)`, since we iterate through the list once.
- **Space Complexity:** `O(1)`, as we use only a single variable (`answer`).

This is a **brilliant and elegant solution** leveraging XOR to efficiently find the unique number! 🚀🔥

