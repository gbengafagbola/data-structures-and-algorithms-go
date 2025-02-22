## Majority Element Algorithm

This repository contains two implementations of the **Majority Element** algorithm in Go:
1. **Hash Map Counting Approach** - Uses a frequency map to count occurrences.
2. **Boyer-Moore Voting Algorithm** - Optimized approach with O(1) space complexity.

## Problem Statement
Given an array `nums` of size `n`, find the majority element. The majority element is the element that appears more than `⌊n / 2⌋` times. You may assume that the majority element always exists.

---
## Approach 1: Hash Map Counting
### Algorithm:
1. Create a hash map (`tab`) to store counts of each element.
2. Iterate through `nums`, incrementing counts in `tab`.
3. If any element's count exceeds `n/2`, return it immediately.

### Code:
```go
func majorityElement(nums []int) int {
    check := len(nums) / 2
    tab := make(map[int]int)

    for _, num := range nums {
        tab[num]++
        if tab[num] > check {
            return num
        }
    }
    return 0
}
```

### Complexity Analysis:
- **Time Complexity:** O(n) (Single pass through the array)
- **Space Complexity:** O(n) (Uses a hash map to store counts)

### Visualization:
#### Example: `nums = [2,2,1,1,1,2,2]`
| Step | Number | Hash Map (`tab`) | Majority Element Found? |
|------|--------|------------------|------------------------|
| 1    | 2      | `{2:1}`           | No |
| 2    | 2      | `{2:2}`           | No |
| 3    | 1      | `{2:2, 1:1}`      | No |
| 4    | 1      | `{2:2, 1:2}`      | No |
| 5    | 1      | `{2:2, 1:3}`      | No |
| 6    | 2      | `{2:3, 1:3}`      | No |
| 7    | 2      | `{2:4, 1:3}`      | Yes (2 is the majority) |

---
## Approach 2: Boyer-Moore Voting Algorithm
### Algorithm:
1. Initialize `majority = nums[0]`, `count = 1`.
2. Iterate through `nums`:
   - If `nums[i] == majority`, increment `count`.
   - Else, decrement `count`.
   - If `count == 0`, update `majority = nums[i]` and reset `count = 1`.
3. Return `majority` at the end.

### Code:
```go
func majorityElement2(nums []int) int {
    count := 1
    majority := nums[0]
    n := len(nums)

    for i := 1; i < n; i++ {
        if majority == nums[i] {
            count++
        } else {
            count--
            if count == 0 {
                majority = nums[i]
                count = 1
            }
        }
    }
    return majority
}
```

### Complexity Analysis:
- **Time Complexity:** O(n) (Single pass through the array)
- **Space Complexity:** O(1) (Uses only two extra variables)

### Visualization:
#### Example: `nums = [2,2,1,1,1,2,2]`
| Step | Number | Current Majority | Count |
|------|--------|-----------------|-------|
| 1    | 2      | 2               | 1     |
| 2    | 2      | 2               | 2     |
| 3    | 1      | 2               | 1     |
| 4    | 1      | 2               | 0     | (Reset Majority)
| 5    | 1      | 1               | 1     |
| 6    | 2      | 1               | 0     | (Reset Majority)
| 7    | 2      | 2               | 1     |
**Output:** `2` (Majority Element)

---
## Comparison of Both Approaches
| Approach                | Time Complexity | Space Complexity | Best for |
|-------------------------|----------------|-----------------|------------|
| Hash Map Counting       | O(n)           | O(n)            | When space is not a constraint |
| Boyer-Moore Voting      | O(n)           | O(1)            | When space optimization is needed |

## Running the Code
To execute the program, run the following in your terminal:
```sh
go run main.go
```

### Example Output:
```sh
2
```

## Conclusion
Both approaches efficiently find the majority element. The **hash map approach** is more intuitive but uses extra space, while the **Boyer-Moore Voting Algorithm** is optimal in terms of space. 🚀

