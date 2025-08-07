Great! Here's an **expanded version** of the `README.md` with detailed **explanation of the code**, **step-by-step sliding mechanism**, **diagrams**, and **comments** that explain **every important line**.

---

# 🔤 `maxVowels` – Maximum Number of Vowels in a Substring of Length `k`

---

## 📚 Problem Statement

You are given a **string** `s` and an **integer** `k`.

Your task is to find the **maximum number of vowels** that appear in any substring of `s` with length exactly `k`.

---

## ✅ What are Vowels?

We only care about lowercase vowels:

```
a, e, i, o, u
```

---

## 🎯 Goal

Return the **maximum number** of vowels found in **any substring of length `k`**.

---

## 🔁 Sliding Window Technique – Visual Breakdown

### Input

```go
s := "abciiidef"
k := 3
```

We slide a **window of size `k = 3`** across the string and count how many vowels are in that window.

### Step-by-step Window Movement

```
Initial:       [a b c] i i i d e f   → 1 vowel (a)
Slide →        a [b c i] i i d e f   → 1 vowel (i)
Slide →        a b [c i i] i d e f   → 2 vowels (i, i)
Slide →        a b c [i i i] d e f   → 3 vowels (i, i, i) ✅ max
Slide →        a b c i [i i d] e f   → 2 vowels (i, i)
Slide →        a b c i i [i d e] f   → 2 vowels (i, e)
Slide →        a b c i i i [d e f]   → 1 vowel (e)
```

✔️ **Maximum number of vowels in any substring of length 3 is 3**

---

## 🔍 Code with Line-by-Line Explanation

```go
// Define a map to check if a character is a vowel in O(1) time
var vowels = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}
```

We use a **map of bytes** to store vowel characters so that `vowels[someChar]` returns `true` only if it's a vowel.

---

```go
func maxVowels(s string, k int) int {
	count := 0
```

`count` keeps track of the number of vowels **in the current window**.

---

```go
	// Count vowels in the first window of size k
	for i := 0; i < k; i++ {
		if vowels[s[i]] {
			count++
		}
	}
```

🚪 **Initialize the window**: Check the first `k` characters of the string and count how many are vowels.

---

```go
	maxCount := count
```

🔒 `maxCount` stores the **maximum number of vowels** seen in any window so far.

---

```go
	// Slide the window across the string
	for i := k; i < len(s); i++ {
```

➡️ We now start sliding the window **one character at a time**.

At each step:

* Add one character (to the right)
* Remove one character (from the left)

---

```go
		if vowels[s[i]] {
			count++ // include the new character on the right
		}
```

➕ Check if the new rightmost character is a vowel. If yes, increment the `count`.

---

```go
		if vowels[s[i-k]] {
			count-- // remove the leftmost character from the window
		}
```

➖ Check if the character that just left the window (i.e., `s[i-k]`) is a vowel. If yes, decrement the `count`.

---

```go
		if count > maxCount {
			maxCount = count
		}
	}
```

🔼 If the current window's vowel count is **greater than** `maxCount`, update it.

---

```go
	return maxCount
}
```

🎉 Finally, return the **maximum number of vowels** found in any window.

---

## 📈 Time and Space Complexity

| Complexity Type | Value                                       |
| --------------- | ------------------------------------------- |
| ⏱ Time          | O(n), where `n` is length of string `s`     |
| 🧠 Space        | O(1), since the vowel map has constant size |

---

## 🧪 Test Cases

```go
fmt.Println(maxVowels("abciiidef", 3)) // 3
fmt.Println(maxVowels("aeiou", 2))    // 2
fmt.Println(maxVowels("leetcode", 3)) // 2
fmt.Println(maxVowels("rhythms", 4))  // 0
fmt.Println(maxVowels("tryhard", 4))  // 1
```

---

## 🧠 How the Sliding Window Works – Animation

```
s:  a  b  c  i  i  i  d  e  f
     ^-----^                (initial window)
        ^-----^             (slide right)
           ^-----^          ...
```

At every step:

* The window shifts one index to the right
* The vowel count is updated in **O(1)** time
* No need to re-scan the entire window

This makes the algorithm highly **efficient** for large strings.

---

Here's your updated `README.md`, now including both the **map-based** and **array-based** sliding window approaches, with a side-by-side comparison, code explanation, and when to prefer each.

---

# 🔤 `maxVowels` – Maximum Number of Vowels in a Substring of Length `k`

---

## 📚 Problem Statement

Given a string `s` and an integer `k`, return the **maximum number of vowels** in any substring of length `k`.

---

## ✅ What Are Vowels?

We are only interested in lowercase English vowels:

```
a, e, i, o, u
```

---

## 🎯 Goal

Find the **maximum number of vowels** in **any substring of length `k`** of the string `s`.

---

## 🚀 Approaches

### 1️⃣ Map-Based Sliding Window (Classic)

```go
var vowels = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

func maxVowels(s string, k int) int {
	count := 0
	for i := 0; i < k; i++ {
		if vowels[s[i]] {
			count++
		}
	}

	maxCount := count

	for i := k; i < len(s); i++ {
		if vowels[s[i]] {
			count++
		}
		if vowels[s[i-k]] {
			count--
		}
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
```

### 2️⃣ Array-Based Sliding Window (Optimized)

```go
func maxVowels(s string, k int) int {
	vowels := make([]int, 26)
	vowels['a'-'a'] = 1
	vowels['e'-'a'] = 1
	vowels['i'-'a'] = 1
	vowels['o'-'a'] = 1
	vowels['u'-'a'] = 1

	count := 0
	for i := 0; i < k; i++ {
		count += vowels[s[i]-'a']
	}
	maxCount := count

	for i := k; i < len(s); i++ {
		count -= vowels[s[i-k]-'a']
		count += vowels[s[i]-'a']
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
```

---

## 🧠 Sliding Window Visual

For input `s = "abciiidef"`, `k = 3`:

```
Window movement:

[ a b c ] i i i d e f   → 1 vowel
  a [ b c i ] i i d e f → 1 vowel
    a b [ c i i ] i d e → 2 vowels
      a b c [ i i i ] d → 3 vowels ✅ max
         ...
```

---

## ⚖️ Comparison

| Feature          | Map-Based Approach         | Array-Based Approach              |
| ---------------- | -------------------------- | --------------------------------- |
| Vowel check      | `vowels[char]`             | `vowels[char - 'a']`              |
| Time Complexity  | O(n)                       | O(n)                              |
| Space Complexity | O(1) (for 5 vowel entries) | O(1) (fixed 26-element int array) |
| Speed            | Slightly slower (hash)     | Slightly faster (index)           |
| Clarity          | More readable              | Slightly more technical           |

---

## 🧪 Sample Test Cases

```go
fmt.Println(maxVowels("abciiidef", 3)) // Output: 3
fmt.Println(maxVowels("aeiou", 2))    // Output: 2
fmt.Println(maxVowels("leetcode", 3)) // Output: 2
fmt.Println(maxVowels("rhythms", 4))  // Output: 0
fmt.Println(maxVowels("tryhard", 4))  // Output: 1
```

---

## 📈 Time and Space Complexity

| Metric           | Value |
| ---------------- | ----- |
| Time Complexity  | O(n)  |
| Space Complexity | O(1)  |

Both versions are optimal in time and space.

---

## 🧠 Summary

* Use the **map-based version** for readability and clarity.
* Use the **array-based version** for **maximum speed** and **lower memory overhead**.
* Both solve the problem in **linear time**, using an efficient **sliding window** technique.

---

Let me know if you’d like this in markdown file format or want to include a benchmark comparison too.
