# Rescue Boats Required

## Problem Statement
Given an array `people` where `people[i]` represents the weight of the `i`-th person and an integer `limit`, representing the maximum weight a boat can carry, determine the minimum number of boats required to rescue all people.

Each boat can carry at most two people at the same time, provided their combined weight does not exceed the `limit`.

---

## Algorithm Explanation

The approach follows a **two-pointer technique**:

1. **Sort the array** in ascending order.
2. Use two pointers:
   - `smallest` (starting from the lightest person, `index 0`)
   - `biggest` (starting from the heaviest person, `index len(people) - 1`)
3. Try to pair the lightest and heaviest person together:
   - If their combined weight is within the limit, move both pointers (`smallest++`, `biggest--`).
   - Otherwise, send the heaviest person alone (`biggest--`).
4. Count the number of boats used.

This ensures the minimum number of boats is used efficiently.

---

## Code Implementation

```go
package main

import (
	"fmt"
	"sort"
)

func rescueBoatsRequired(people []int, limit int) int {
	sort.Ints(people)

	smallest := 0
	biggest := len(people) - 1
	boats := 0

	for smallest <= biggest {
		if people[smallest]+people[biggest] <= limit {
			smallest++
			biggest--
		} else {
			biggest--
		}
		boats++
	}

	return boats
}

func main() {
	people := []int{3, 2, 4, 5, 1}
	limit := 6

	result := rescueBoatsRequired(people, limit)
	fmt.Println(result) // Output: 3
}
```

---

## Example Walkthrough
### Input:
```go
people = [3, 2, 4, 5, 1]
limit = 6
```

### Sorting:
```go
[1, 2, 3, 4, 5]
```

### Step-by-step Execution:
| Boat | Pair Chosen | Remaining People |
|------|------------|------------------|
| 1    | (1, 5)     | [2, 3, 4]        |
| 2    | (2, 4)     | [3]              |
| 3    | (3)        | []               |

✅ **Output: `3` boats required**

---

## Complexity Analysis
- **Sorting takes** `O(n log n)`
- **Two-pointer traversal takes** `O(n)`
- **Total Complexity:** `O(n log n)`
- **Space Complexity:** `O(1)`, as sorting is done in-place.

---

## Edge Cases Considered
| Case | Example Input | Expected Output |
|------|--------------|----------------|
| All people must go alone | `[4, 4, 4, 4]`, limit = `4` | `4` |
| Optimally paired people | `[1, 2, 3, 4, 5]`, limit = `6` | `3` |
| Single person | `[5]`, limit = `6` | `1` |

---

## Summary
This implementation effectively finds the minimum number of boats using sorting and a **greedy** two-pointer approach. It ensures that heavier individuals are accommodated efficiently while minimizing the total boat usage.

