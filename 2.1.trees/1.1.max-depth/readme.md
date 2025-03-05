# **Find the Maximum Depth of a Binary Tree**  

## **Introduction**  
This program calculates the **maximum depth (height) of a binary tree**, which is the longest path from the **root to a leaf node**.

---

## **What is the Maximum Depth of a Tree?**  

The **maximum depth** (or height) of a binary tree is the **number of edges** on the longest path from the root to a leaf.

✅ **Example 1:**  
```
        1
       / \
      2   3
     / \
    4   5
```
📌 **Max Depth = 3**  
(Path: `1 → 2 → 4` or `1 → 2 → 5`)

✅ **Example 2:**  
```
    1
   /
  2
 / 
3
```
📌 **Max Depth = 3**  
(Path: `1 → 2 → 3`)

---

## **How the Code Works**
### **1. Struct Definition (`TreeNode`)**
We define a tree node structure:

```go
type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}
```
Each node contains:
- `Val` → The node’s value
- `Left` → Pointer to the left child
- `Right` → Pointer to the right child

---

### **2. Recursive Function `maxDepth(root *TreeNode) int`**  
This function **recursively calculates** the depth of a tree.

```go
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0  // An empty tree has depth 0
	}

	left := 1 + maxDepth(root.Left)  // Calculate left subtree depth
	right := 1 + maxDepth(root.Right) // Calculate right subtree depth

	ans := max(left, right)  // Choose the maximum depth
	return ans
}
```

🔹 **Key Observations:**
- If `root == nil`, the tree is **empty** → depth = `0`.
- Recursively **calculate depth** for left and right subtrees.
- The final depth is **`1 + max(left, right)`**, adding **1** for the current node.

📌 **Illustration of Recursion:**  
For this tree:
```
       1
      / \
     2   3
    /
   4
```
1. `maxDepth(4) = 1` (leaf node)
2. `maxDepth(2) = max(1, 0) + 1 = 2`
3. `maxDepth(3) = 1` (leaf node)
4. `maxDepth(1) = max(2, 1) + 1 = 3`

---

### **3. `max(a, b) int` (Helper Function)**
This function returns the maximum of two numbers:

```go
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

---

## **Example Walkthrough**
Let’s analyze a test case:

```go
root := &TreeNode{1,
    &TreeNode{2, &TreeNode{4, nil, nil}, nil},
    &TreeNode{3, nil, nil},
}
fmt.Println(maxDepth(root)) // Output: 3
```
### **Step-by-Step Execution**
1. **Compute `maxDepth(4)`**:
   ```
   root = 4 → Left: nil, Right: nil
   return 1 (leaf node)
   ```
2. **Compute `maxDepth(2)`**:
   ```
   maxDepth(Left: 4) = 1
   maxDepth(Right: nil) = 0
   maxDepth(2) = max(1, 0) + 1 = 2
   ```
3. **Compute `maxDepth(3)`**:
   ```
   maxDepth(3) = 1 (leaf node)
   ```
4. **Compute `maxDepth(1)`**:
   ```
   maxDepth(Left: 2) = 2
   maxDepth(Right: 3) = 1
   maxDepth(1) = max(2, 1) + 1 = 3
   ```

✅ **Final Answer: `3`**

---

## **Final Thoughts**
✅ **Time Complexity:** **O(n)** (we visit each node once)  
✅ **Space Complexity:** **O(n)** (recursive call stack)  

🔹 **Summary**
- **Recursively compute left & right subtree depths.**
- **Take the maximum depth and add 1 (for the root).**
- Efficient **O(n) solution** using **DFS recursion**.

---

Would you like an **iterative solution using BFS (level order traversal)?** 🚀