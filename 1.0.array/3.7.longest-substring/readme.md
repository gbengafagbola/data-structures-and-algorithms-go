# Longest Substring Without Repeating Characters

## Overview
This Go program finds the length of the longest substring without repeating characters in a given string. The implementation uses the sliding window technique with a hash map to efficiently track character positions and dynamically adjust the window.

## Code
```go
package main

import (
	"fmt"
	"math"
)

func lengthOfLoongestSubString(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	left, right := 0, 0
	maxLen := 0

	chars := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if val, ok := chars[s[i]]; ok {
			left = int(math.Max(float64(left), float64(val+1)))
		}
		right++
		chars[s[i]] = i
		maxLen = int(math.Max(float64(right-left), float64(maxLen)))
	}

	return maxLen
}

func main() {
	s := "abcabcbb"
	result := lengthOfLoongestSubString(s)
	fmt.Println(result)
}
```

## Flow Breakdown
### 1. **Initialization**
- A hash map `chars` stores the last seen index of characters.
- `left` and `right` pointers define the sliding window.
- `maxLen` keeps track of the maximum substring length.

### 2. **Sliding Window Iteration**
- Iterate through the string character by character.
- If the character is found in `chars`, move `left` to `max(last_seen_index+1, left)`, ensuring unique characters within the window.
- Update the `chars` map with the current character's index.
- Update `maxLen` by calculating the window size (`right - left`).

### 3. **Result Computation**
- At each step, the maximum length of a substring without repeating characters is stored in `maxLen`.
- Finally, the function returns `maxLen`.

## Visual Flow
```
Input: "abcabcbb"

Step-by-step tracking:

1. [a] -> maxLen = 1
2. [ab] -> maxLen = 2
3. [abc] -> maxLen = 3
4. [bca] -> maxLen = 3 (shift left pointer)
5. [cab] -> maxLen = 3
6. [abc] -> maxLen = 3
7. [b] -> maxLen = 3 (duplicate 'b' forces reset)
8. [bb] -> maxLen = 3 (reset again)

Output: 3
```

## Fixing Issues
- There's a typo in the function name: `lengthOfLoongestSubString` should be `lengthOfLongestSubString`.
- Instead of incrementing `right`, the loop index `i` already represents `right`, making it redundant.

## Optimized Code
```go
func lengthOfLongestSubString(s string) int {
	left, maxLen := 0, 0
	chars := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		if val, ok := chars[s[right]]; ok {
			left = int(math.Max(float64(left), float64(val+1)))
		}
		chars[s[right]] = right
		maxLen = int(math.Max(float64(right-left+1), float64(maxLen)))
	}

	return maxLen
}
```

## Conclusion
This implementation efficiently finds the longest substring without repeating characters using the sliding window technique and a hash map. With an optimized approach, it runs in **O(n)** time complexity.

 