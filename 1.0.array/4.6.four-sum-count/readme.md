# Four Sum Count Algorithm Breakdown

## Problem Statement
Given four integer arrays `A`, `B`, `C`, and `D`, each of length `n`, compute how many tuples `(i, j, k, l)` exist such that:

\[ A[i] + B[j] + C[k] + D[l] = 0 \]

## Algorithm Explanation
The function `fourSumCount` efficiently solves this problem using a **hash map (dictionary)** to store intermediate sums, which reduces the time complexity compared to a brute-force approach.

## Approach 1: Basic Hash Map Optimization

### Steps Breakdown
### Step 1: Precompute Partial Sums for `A` and `B`
We iterate over every combination of elements from `A` and `B` and store the sum frequencies in a hash map `m`.

#### Example
Let’s assume:

```
A = [1, 2]
B = [-2, -1]
C = [-1, 2]
D = [0, 2]
```

We compute:

```
(1) + (-2) = -1
(1) + (-1) =  0
(2) + (-2) =  0
(2) + (-1) =  1
```

So, the hash map `m` stores:

```
m = {
    -1: 1,  // Sum -1 appears once
     0: 2,  // Sum  0 appears twice
     1: 1   // Sum  1 appears once
}
```

### Step 2: Compute Complementary Sums for `C` and `D`
Now, we iterate over every pair from `C` and `D`, compute their sum, negate it, and check if it exists in `m`.

For our example:

```
(-1) + (0)  = -1 → Complement: 1
(-1) + (2)  =  1 → Complement: -1
(2)  + (0)  =  2 → Complement: -2
(2)  + (2)  =  4 → Complement: -4
```

We check these values in `m`:

- Complement `1` exists in `m`, so we add `1` (count from `m[1]`).
- Complement `-1` exists in `m`, so we add `1` (count from `m[-1]`).

Total valid tuples: `2`.

### Code
```go
func fourSumCount(A []int, B []int, C []int, D []int) int {
    m := map[int]int{}

    // Step 1: Compute sums of pairs from A and B and store in map
    for _, a := range A {
        for _, b := range B {
            target := a + b
            if _, ok := m[target]; !ok {
                m[target] = 0
            }
            m[target]++
        }
    }

    ans := 0
    
    // Step 2: Compute sums of pairs from C and D and check in map
    for _, c := range C {
        for _, d := range D {
            target := -(c + d)
            if _, ok := m[target]; ok {
                ans += m[target]
            }
        }
    }
    return ans
}
```

## Approach 2: Optimized Hash Map with Preallocation
This approach improves performance by:
1. **Using a preallocated map** to reduce memory allocation overhead.
2. **Directly incrementing values in the map** to minimize conditional checks.

### Code
```go
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    data := make(map[int]int, len(nums1)*len(nums2))
    for _, num1 := range nums1 {
        for _, num2 := range nums2 {
            data[num1+num2]++
        }
    }
    var ret int
    for _, num3 := range nums3 {
        for _, num4 := range nums4 {
            ret += data[-num3-num4]
        }
    }
    return ret
}
```

### Why Preallocate `data`?
The line:
```go
data := make(map[int]int, len(nums1)*len(nums2))
```
**Optimizes performance by:**
- **Avoiding unnecessary map resizing**: Go dynamically resizes maps, which incurs overhead.
- **Reducing memory fragmentation**: Allocating upfront ensures contiguous memory usage.
- **Speeding up insertions**: Directly inserting values without checking conditions improves efficiency.

Since `nums1` and `nums2` contribute `len(nums1) * len(nums2)` sums, preallocating ensures the map has enough capacity.

## Complexity Analysis
- **Step 1:** Constructing the hash map takes **O(n²)** time.
- **Step 2:** Checking complementary sums takes **O(n²)** time.
- **Total complexity:** **O(n²) + O(n²) = O(n²)** (efficient compared to O(n⁴)).

## Edge Cases Considered
1. **All elements are zero** → The number of valid tuples is maximum.
2. **All elements are distinct** → The number of valid tuples could be zero.
3. **Large values** → The function handles them efficiently using a hash map.

## Summary
- **Uses a hash map to store intermediate sums**
- **Reduces complexity from O(n⁴) to O(n²)**
- **Efficiently finds all valid `(i, j, k, l)` tuples**

