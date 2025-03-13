Let's break it down step by step:

---

### **1. Struct Definition**
```go
type replacement struct {
	index  int
	Before string
	After  string
}
```
- This struct represents a **text replacement rule**.
- `index`: The position in the string where we want to check for the replacement.
- `Before`: The substring that we want to replace.
- `After`: The new substring that will replace `Before`.

---

### **2. The `replace` Method**
```go
func (r replacement) replace(sourceCode string) string {
	if strings.HasPrefix(sourceCode[r.index:], r.Before) {
		return sourceCode[:r.index] + r.After + sourceCode[r.index+len(r.Before):]
	}
	return sourceCode
}
```

#### **How It Works:**
1. **`strings.HasPrefix(sourceCode[r.index:], r.Before)`**
   - This checks if the substring starting from `index` matches `Before`.
   - If `index = 0` and `Before = "Hi"`, then `sourceCode[0:] = "Hi There"`, so the function checks if `"Hi There"` starts with `"Hi"`.

2. **If the condition is true:**
   - `sourceCode[:r.index]`: This takes the part **before** the `index`.  
   - `r.After`: This is the new replacement string.
   - `sourceCode[r.index+len(r.Before):]`: This takes the part **after** the `Before` string.
   - The result is: `Before` is replaced with `After` at the specified `index`.

3. **If the condition is false:**  
   - The original string is returned as no replacement is needed.

---

### **3. `main` Function**
```go
func main() {
	sourceCode := "Hi There"
	r := replacement{index: 0, Before: "Hi", After: "Hello"}
	newCode := r.replace(sourceCode)
	fmt.Println(newCode) // Output: "Hello There"
}
```
- Creates a string `sourceCode = "Hi There"`.
- Defines a `replacement` rule: replace `"Hi"` at index `0` with `"Hello"`.
- Calls `replace(sourceCode)`, which modifies and returns the new string.
- Prints the output:  
  ```
  Hello There
  ```

---

### **Key Takeaways**
✅ **Uses string slicing to modify the string efficiently.**  
✅ **Ensures replacement happens only at the correct index.**  
✅ **Returns a new string because strings are immutable in Go.**  