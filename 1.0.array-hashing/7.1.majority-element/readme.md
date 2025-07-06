Sure! Here's a `README.md` that thoroughly explains your `MajorityElement` program, including logic, breakdown, and example usage:

---

# MajorityElement (Go)

This Go program identifies the **majority element** in an integer array—i.e., the element that appears **more than n/2 times**, where `n` is the length of the array.

If such an element exists, it returns the majority element. Otherwise, it returns `-1`.

---

## 📦 Function Signature

```go
func MajorityElement(array []int) int
```

---

## ✅ Example

```go
Input:  [2, 2, 1, 2, 3, 2, 2]
Output: 2
```

---

## 🧠 How It Works

### 1. Frequency Map Construction

The function first builds a map to count the number of times each integer appears.

```go
n := make(map[int]int)

for _, num := range array {
    if _, seen := n[num]; seen {
        n[num] += 1
    } else {
        n[num] = 1
    }
}
```

### 2. Extract All Occurrence Values

We extract all the frequency counts from the map and store them in a slice.

```go
occurence := []int{}
for _, value := range n {
    occurence = append(occurence, value)
}
```

### 3. Sort Occurrences to Get Max Frequency

By sorting the slice of occurrences, we can easily identify the maximum value (most frequent occurrence).

```go
sort.Ints(occurence)
v := occurence[len(occurence)-1]
```

### 4. Validate and Return the Majority Element

We check whether this maximum frequency is greater than `n/2`. If so, we loop through the map to find the key (i.e., the number) that has this frequency.

```go
if v > len(array)/2 {
    for key, value := range n {
        if value == v {
            return key
        }
    }
}
```

If no such element exists, return `-1`.

---

## ⚠️ Edge Cases

* If the array is empty → returns `-1`.
* If there's no element that appears more than `n/2` times → returns `-1`.

---

## 🛠️ Time and Space Complexity

* **Time Complexity**: `O(n log n)`

  * Building the frequency map: `O(n)`
  * Sorting the frequencies: `O(n log n)`

* **Space Complexity**: `O(n)`

  * Due to the frequency map and the slice of occurrence counts.

---

## 🔁 Example Usage

```go
package main

import "fmt"

func main() {
	arr := []int{3, 3, 4, 2, 3, 3, 3}
	fmt.Println(MajorityElement(arr)) // Output: 3
}
```

---

## 📌 Notes

While this solution works, it's not optimal in terms of space or time due to sorting. For a more efficient solution, consider **Moore’s Voting Algorithm**, which works in `O(n)` time and `O(1)` space.

Let me know if you’d like to see that version too!

---
