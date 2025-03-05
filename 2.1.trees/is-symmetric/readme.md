# **Check if a Binary Tree is Symmetric (Mirror Image of Itself)**  

## **Introduction**  
This program determines whether a binary tree is **symmetric**, meaning it is a **mirror image** of itself along its center.

---

## **What is a Symmetric Tree?**  

A tree is symmetric if the left and right subtrees are mirror images of each other.

✅ **Example of a Symmetric Tree:**  
```
        1
       / \
      2   2
     / \ / \
    3  4 4  3
```
This tree is **mirrored** along the center.

❌ **Example of a Non-Symmetric Tree:**  
```
        1
       / \
      2   2
       \    \
       3     3
```
This tree **is not symmetric** because the left subtree has `3` on the right, while the right subtree has `3` on the right instead of the left.

---

## **How the Code Works**
### **1. Struct Definition (`TreeNode`)**
We define a tree node structure:
```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
```
Each node contains:
- `Val` → The node’s value
- `Left` → Pointer to the left child
- `Right` → Pointer to the right child

---

### **2. `isMirror(t1 *TreeNode, t2 *TreeNode) bool`**  
This function **checks if two trees are mirror images** of each other.

```go
func isMirror(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true  // Both subtrees are empty -> Symmetric
	}

	if t1 == nil || t2 == nil {
		return false // One subtree is empty, but the other is not -> Not symmetric
	}

	// Check if values match and left-right / right-left subtrees are mirror images
	return t1.Val == t2.Val && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}
```
🔹 **Key Observations:**
- If both `t1` and `t2` are `nil`, they are symmetric.
- If only one is `nil`, it’s **not symmetric**.
- Otherwise, we **recursively check**:
  - `t1.Left` vs `t2.Right`
  - `t1.Right` vs `t2.Left`

📌 **Illustration of Recursion:**  
For this tree:
```
       1
      / \
     2   2
    / \ / \
   3  4 4  3
```
- `isMirror(2, 2)` → ✅ (values match)
- `isMirror(3, 3)` → ✅ (left subtree of left == right subtree of right)
- `isMirror(4, 4)` → ✅ (right subtree of left == left subtree of right)

Since all **subtree comparisons pass**, the function **returns `true`**.

---

### **3. `isSymmetric(root *TreeNode) bool`**  
This function starts the symmetry check from the root.

```go
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true // An empty tree is symmetric
	}

	return isMirror(root.Left, root.Right)
}
```
- If `root == nil`, the tree is empty → **always symmetric**.
- Otherwise, we check if the **left and right subtrees** are mirrors.

---

## **Example Walkthrough**
Let’s analyze a test case:

```go
root := &TreeNode{1,
    &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
    &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}},
}
fmt.Println(isSymmetric(root)) // Output: true
```
### **Step-by-Step Execution**
1. Check `root.Left (2)` vs `root.Right (2)` → ✅ (values match)
2. Check `root.Left.Left (3)` vs `root.Right.Right (3)` → ✅
3. Check `root.Left.Right (4)` vs `root.Right.Left (4)` → ✅
4. No more nodes left to check → Return **`true`** ✅

---

## **Final Thoughts**
✅ **Time Complexity:** **O(n)** (we visit each node once)  
✅ **Space Complexity:** **O(n)** (for recursive stack in worst case)  

🔹 **Summary**
- **Recursively** check if left & right subtrees are **mirrors**.
- **Base cases:** If both `nil` → ✅, if only one `nil` → ❌.
- Efficient **O(n) solution** using **DFS recursion**.

Would you like an **iterative solution** using a queue? 🚀