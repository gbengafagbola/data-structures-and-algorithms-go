# Non-Constructible Change Problem

## Problem Statement
Given an array of positive integers representing coin denominations, return the minimum amount of change that **cannot** be created. You can assume you have an infinite number of each coin.

---

## Greedy Solution Explanation

### Idea
If you can create all the change from `1` to `x`, and the next coin is greater than `x + 1`, then you can't create `x + 1`. That value is the smallest non-constructible change.

### Step-by-Step Example

#### Input:
```go
coins := []int{5, 7, 1, 1, 2, 3, 22}
```

#### Step 1: Sort the coins
```
[1, 1, 2, 3, 5, 7, 22]
```

#### Step 2: Initialize `currentChangeCreated = 0`

#### Step 3: Walk through each coin
| Coin | Can we create `coin > currentChangeCreated + 1`? | `currentChangeCreated` After Adding |
|------|---------------------------------------------------|-----------------------------------|
| 1    | 1 <= 0 + 1 ✅                                     | 1                                 |
| 1    | 1 <= 1 + 1 ✅                                     | 2                                 |
| 2    | 2 <= 2 + 1 ✅                                     | 4                                 |
| 3    | 3 <= 4 + 1 ✅                                     | 7                                 |
| 5    | 5 <= 7 + 1 ✅                                     | 12                                |
| 7    | 7 <= 12 + 1 ✅                                    | 19                                |
| 22   | 22 > 19 + 1 ❌                                   | Return 20                         |

---

## Final Answer
```
20
```

You can construct every value from 1 to 19, but **not 20**.

---

## Time Complexity
- **Sorting:** O(n log n)
- **Loop:** O(n)
- **Total:** O(n log n)

This is far more efficient than checking all subsets (which is exponential in time).

---

## Summary
This greedy approach is optimal because it only cares about whether we can extend the range of constructible change. The moment we hit a gap, that's our answer.

