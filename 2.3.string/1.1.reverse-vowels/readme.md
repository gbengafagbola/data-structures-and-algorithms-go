Sure! Here's a concise and clear `README.md` for your `reverseVowels` function:

---

# Reverse Vowels in a String – Go Implementation

This Go program reverses only the vowels in a given string, preserving the positions of all other characters.

## 🧠 Problem Statement

Given a string `s`, return a new string where **only the vowels** (`a`, `e`, `i`, `o`, `u` in both lowercase and uppercase) are reversed. The positions of all consonants and other characters remain unchanged.

### ✍️ Example

```go
Input:  "hello"
Output: "holle"

Input:  "leetcode"
Output: "leotcede"
```

---

## 🚀 How It Works

* Two-pointer approach:

  * Start with one pointer `i` from the beginning and another `j` from the end.
  * Move both pointers toward the center:

    * If both characters are vowels, swap them.
    * If either one is not a vowel, move the respective pointer accordingly.
* Convert the string to a slice of runes to allow in-place modification and handle Unicode characters correctly.

---

## 🧩 Functions

### `reverseVowels(s string) string`

Main function that:

* Converts string to `[]rune`.
* Uses a two-pointer technique to reverse only the vowels.
* Returns the final string after swapping.

### `isVowel(vowel rune) bool`

Utility function that:

* Returns `true` if the character is a vowel (`a, e, i, o, u` in upper or lower case).
* Returns `false` otherwise.

---

## 🧪 Example Code

```go
func main() {
    fmt.Println(reverseVowels("hello"))      // Output: "holle"
    fmt.Println(reverseVowels("leetcode"))   // Output: "leotcede"
}
```

---

## 📦 Run Locally

1. Save the code in a `.go` file, for example, `reverse_vowels.go`.
2. Run using:

```bash
go run reverse_vowels.go
```

---

## ✅ Notes

* This implementation assumes ASCII input.
* For full Unicode support (e.g., accented vowels), the `isVowel` function would need to be extended accordingly.

