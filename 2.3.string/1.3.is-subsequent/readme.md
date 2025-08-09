# 📘 isSubsequence - Explained Like You're 10

## 🧠 What does this do?

Imagine you have **two strings** (just a bunch of letters).

We want to **check** if the first string (`s`) can be **found inside** the second string (`t`) — **in the same order**, but **not necessarily all together**.

### 💡 Example:

```go
s = "abc"
t = "ahbgdc"
```

We are asking: "Can we find the letters `a`, `b`, and `c` in `t` in the same order as in `s`?"

➡️ Yes! Look:

* `a` is there ✅
* then `h` (we skip it)
* then `b` ✅
* then `g` (skip)
* then `d` (skip)
* then `c` ✅

So `s` is a **subsequence** of `t`.

---

## 👨‍💻 What does the code do?

Here's the code:

```go
func isSubsequence(s string, t string) bool {
    i, j := 0, 0

    for i < len(s) && j < len(t) {
        if s[i] == t[j] {
            i++  // found a letter from s in t, move to next letter in s
        }
        j++  // always move to next letter in t
    }

    return i == len(s)  // if we checked all letters in s, return true!
}
```

---

## 🧒 How to Think About It

It's like playing a game where you try to **spot each letter of your secret word (`s`)** inside a long story (`t`) — in the **same order**.

If you can find them all, you win! 🎉

---

## ✅ When does it say "Yes"?

* If you can find **all letters of `s`** inside `t` in order.

## ❌ When does it say "No"?

* If you **miss any letter**, or the order gets **messed up**.

---

## 🧪 Try It Yourself!

Try with:

```go
s = "ace"
t = "abcde"
```

👉 Yes, it's a subsequence!

Or:

```go
s = "aec"
t = "abcde"
```

👉 No, the `e` comes **after** the `c`, which is **not allowed**.

