## **Challenge: Top K Frequent Elements**

## **Problem**  
Given an array of integers, return the `k` most frequent elements.

---

## **Approach**  

### **1. Count Element Frequencies**
- Use a **map** to store the frequency of each element in the array.
- Each key in the map represents an integer, and its value is the count of occurrences of that integer.

### **2. Use a Bucket Sort Approach**
- Create a **slice of lists** where the index represents the frequency of numbers.
- Each index `i` in this slice contains the numbers that appear exactly `i` times.

### **3. Extract the Top K Frequent Elements**
- Iterate from the **highest frequency down** to collect the most frequent `k` elements.

---

## **Complexity Analysis**
- **Time Complexity**:  
  - Counting frequencies: **O(n)**  
  - Bucket sorting: **O(n)**  
  - Extracting top `k` elements: **O(n)**  
  - **Overall: O(n)** (optimal compared to sorting, which is O(n log n))

- **Space Complexity**:  
  - **O(n)** (for the frequency map and result storage)

---

## **Example Execution**
### **Input**
```go
nums := []int{1, 1, 1, 2, 2, 3}
k := 2
```

### **Step-by-Step Breakdown**
#### **Step 1: Count Occurrences**
```
count = { 1: 3, 2: 2, 3: 1 }
```

#### **Step 2: Bucket Sort into Frequency Array**
```
freq = [[], [3], [2], [1], [], [], []]
         0    1    2    3   4   5   6  (Indexes)
```

#### **Step 3: Collect Top `k` Frequent Elements**
```
Start from highest frequency:
  - freq[3] → Add 1 → res = [1]
  - freq[2] → Add 2 → res = [1, 2]
Stop (we have k = 2 elements)
```

### **Final Output**
```go
Output: [1, 2]
```

---

## **Summary**
✅ **Efficient**: Uses **bucket sort** instead of sorting, keeping it **O(n) time complexity**.  
✅ **Optimized for large datasets** compared to sorting-based approaches.  

[View the problem on LeetCode](https://leetcode.com/problems/top-k-elements-in-list/)  
