Let's break this down step by step with an easy-to-understand approach, using visuals and examples.

---

## **Understanding the Problem**
The function `hasPathSum` checks if there is a path from the **root** of a binary tree to a **leaf node** where the sum of the node values along the path equals a given number.

### **What is a Binary Tree?**
A binary tree is a structure where each node has:
- A **value** (`Val`)
- A **left child** (another TreeNode or `nil`)
- A **right child** (another TreeNode or `nil`)

Example:

```
        5
       / \
      4   8
     /   / \
    11  13  4
   /  \      \
  7    2      1
```

If the given sum is **22**, the function should return **true** because there is a path **5 → 4 → 11 → 2** that sums to 22.

---

## **Step-by-Step Execution**
Let's break down how the function works.

### **Step 1: Function Calls**
1. The function `hasPathSum(root, sum)` starts from the root and calls `hasSum(root, sum, 0)`.
2. `cur` keeps track of the **current sum** along the path.

---

### **Step 2: Recursive Traversal**
The function follows these steps:
- Add the current node’s value to `cur`.
- If this is a **leaf node** (no left or right child), check if `cur` matches `sum`.
- If not, recursively call `hasSum` on the **left and right children**.

---

### **Example Execution**
Given the tree:

```
        5
       / \
      4   8
     /   / \
    11  13  4
   /  \      \
  7    2      1
```

**Checking if there is a path sum of 22:**
1. Start at **5**, current sum = **5**.
2. Move left to **4**, current sum = **5 + 4 = 9**.
3. Move left to **11**, current sum = **9 + 11 = 20**.
4. Move left to **7**, current sum = **20 + 7 = 27** (too large, backtrack).
5. Move right to **2**, current sum = **20 + 2 = 22** (Match! Return `true`).

### **Recursive Calls Breakdown**
```
hasSum(5, 22, 0)  →  hasSum(4, 22, 5)  →  hasSum(11, 22, 9)  
   → hasSum(7, 22, 20) (false)
   → hasSum(2, 22, 20) (true)
```

---

## **Edge Cases**
1. **Empty Tree (`root = nil`)** → Return `false`.
2. **Single Node Tree (`root = 5, sum = 5`)** → Return `true` if root value matches `sum`.
3. **No Valid Path (`sum` is too high or no path matches)** → Return `false`.

---

## **Final Thoughts**
This recursive approach efficiently checks all possible paths using **depth-first search (DFS)**. The function ensures we only check root-to-leaf paths and backtracks when necessary.

Would you like me to modify the function for better readability or optimization? 🚀