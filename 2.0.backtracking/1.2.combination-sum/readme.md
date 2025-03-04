### **README: Combination Sum Problem Using Backtracking**

This document explains the **Combination Sum** problem and how it is solved using **backtracking** in Go. The goal is to find all unique combinations of numbers from a given list (`candidates`) that add up to a specific `target` sum. Each number in `candidates` can be used **multiple times** in the combinations.

---

### **Problem Statement**
Given:
- A list of integers (`candidates`).
- A target integer (`target`).

Find:
- All **unique combinations** of numbers from `candidates` that sum up to `target`.

#### Example:
- Input: `candidates = [2, 3, 6, 7]`, `target = 7`
- Output:
  ```
  [
    [2, 2, 3],
    [7]
  ]
  ```

---

### **Backtracking Approach**
Backtracking is used to explore all possible combinations. Here’s how it works:

1. **Start with an empty combination**.
2. **Add a candidate** to the combination and check if the sum matches the target.
3. If the sum **matches**, add the combination to the result.
4. If the sum is **less than the target**, recursively add more candidates.
5. If the sum **exceeds the target**, backtrack by removing the last candidate and try the next one.

---

### **Code Breakdown**

#### **1. `func combinationSum(candidates []int, target int) [][]int {`**
- This is the main function that initializes the backtracking process.
- It takes:
  - `candidates`: The list of numbers to choose from.
  - `target`: The desired sum.

#### **2. `ans := [][]int{}`**
- This initializes an empty slice (`ans`) to store all valid combinations.

#### **3. `cur := []int{}`**
- This initializes an empty slice (`cur`) to represent the current combination being built.

#### **4. `solution(candidates, &ans, &cur, target, 0, 0)`**
- This calls the recursive `solution` function to start building combinations.

#### **5. `return ans`**
- This returns the final list of valid combinations.

---

#### **6. `func solution(candidates []int, ans *[][]int, cur *[]int, target int, index int, sum int) {`**
- This is the recursive function that implements backtracking.
- It takes:
  - `candidates`: The list of numbers to choose from.
  - `ans`: A pointer to the slice storing all valid combinations.
  - `cur`: A pointer to the current combination being built.
  - `target`: The desired sum.
  - `index`: The current position in `candidates` to start choosing from.
  - `sum`: The current sum of the numbers in `cur`.

---

#### **7. `if sum == target {`**
- This checks if the current combination (`cur`) sums up to the target.
- If it does, the combination is added to `ans`.

#### **8. `*ans = append(*ans, append([]int{}, *cur...))`**
- This adds a **copy** of the current combination (`cur`) to `ans`.
- We use `append([]int{}, *cur...)` to ensure that `ans` stores independent slices.

#### **9. `else if sum < target {`**
- This checks if the current sum is less than the target.
- If it is, we continue adding more numbers to the combination.

#### **10. `n := len(candidates)`**
- This stores the length of the `candidates` list.

#### **11. `for i := index; i < n; i++ {`**
- This loops through the `candidates` starting from the current `index`.
- It ensures that we don’t reuse the same candidate in a way that creates duplicate combinations.

#### **12. `*cur = append(*cur, candidates[i])`**
- This adds the current candidate (`candidates[i]`) to the combination (`cur`).

#### **13. `solution(candidates, ans, cur, target, i, sum+candidates[i])`**
- This is the **recursive call**.
- It explores all combinations that include the current candidate.
- The `index` is set to `i` to allow reusing the same candidate multiple times.

#### **14. `*cur = (*cur)[:len(*cur)-1]`**
- This is the **backtracking step**.
- It removes the last candidate from `cur` so we can try the next candidate.

---

### **Visual Walkthrough**

Let’s walk through an example with `candidates = [2, 3, 6, 7]` and `target = 7`.

#### Step 1: Start with `cur = []`, `sum = 0`
- `index = 0`

#### Step 2: Loop Through `candidates`
1. **First Iteration: `i = 0`, `candidates[i] = 2`**
   - Add `2` to `cur` → `cur = [2]`, `sum = 2`.
   - Recursively call `solution`.

2. **Second Level Recursion**
   - `index = 0` (reuse `2`).
   - Add `2` to `cur` → `cur = [2, 2]`, `sum = 4`.
   - Recursively call `solution`.

3. **Third Level Recursion**
   - `index = 0` (reuse `2`).
   - Add `2` to `cur` → `cur = [2, 2, 2]`, `sum = 6`.
   - Recursively call `solution`.

4. **Fourth Level Recursion**
   - `index = 0` (reuse `2`).
   - Add `2` to `cur` → `cur = [2, 2, 2, 2]`, `sum = 8`.
   - Since `sum > target`, backtrack by removing the last `2` → `cur = [2, 2, 2]`.

5. **Third Level Recursion (continued)**
   - Try the next candidate (`3`).
   - Add `3` to `cur` → `cur = [2, 2, 3]`, `sum = 7`.
   - Since `sum == target`, add `[2, 2, 3]` to `ans`.
   - Backtrack by removing `3` → `cur = [2, 2]`.

6. **Second Level Recursion (continued)**
   - Try the next candidate (`3`).
   - Add `3` to `cur` → `cur = [2, 3]`, `sum = 5`.
   - Recursively call `solution`.

7. **Third Level Recursion**
   - `index = 1` (reuse `3`).
   - Add `3` to `cur` → `cur = [2, 3, 3]`, `sum = 8`.
   - Since `sum > target`, backtrack by removing the last `3` → `cur = [2, 3]`.

8. **Second Level Recursion (continued)**
   - Try the next candidate (`6`).
   - Add `6` to `cur` → `cur = [2, 6]`, `sum = 8`.
   - Since `sum > target`, backtrack by removing `6` → `cur = [2]`.

9. **First Level Recursion (continued)**
   - Try the next candidate (`3`).
   - Add `3` to `cur` → `cur = [3]`, `sum = 3`.
   - Recursively call `solution`.

10. **Second Level Recursion**
    - `index = 1` (reuse `3`).
    - Add `3` to `cur` → `cur = [3, 3]`, `sum = 6`.
    - Recursively call `solution`.

11. **Third Level Recursion**
    - `index = 1` (reuse `3`).
    - Add `3` to `cur` → `cur = [3, 3, 3]`, `sum = 9`.
    - Since `sum > target`, backtrack by removing the last `3` → `cur = [3, 3]`.

12. **Second Level Recursion (continued)**
    - Try the next candidate (`6`).
    - Add `6` to `cur` → `cur = [3, 6]`, `sum = 9`.
    - Since `sum > target`, backtrack by removing `6` → `cur = [3]`.

13. **First Level Recursion (continued)**
    - Try the next candidate (`6`).
    - Add `6` to `cur` → `cur = [6]`, `sum = 6`.
    - Recursively call `solution`.

14. **Second Level Recursion**
    - `index = 2` (reuse `6`).
    - Add `6` to `cur` → `cur = [6, 6]`, `sum = 12`.
    - Since `sum > target`, backtrack by removing the last `6` → `cur = [6]`.

15. **First Level Recursion (continued)**
    - Try the next candidate (`7`).
    - Add `7` to `cur` → `cur = [7]`, `sum = 7`.
    - Since `sum == target`, add `[7]` to `ans`.
    - Backtrack by removing `7` → `cur = []`.

#### Final Result:
```
ans = [
  [2, 2, 3],
  [7]
]
```

---

### **Key Takeaways**
1. Backtracking explores all possible combinations by **building and undoing choices**.
2. The base case ensures we stop when the target is reached.
3. The recursive call allows reusing the same candidate multiple times.

Let me know if you need further clarification! 😊