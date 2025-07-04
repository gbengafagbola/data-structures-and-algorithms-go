package main

func ZeroSumSubarray(nums []int) bool {
    seen := make(map[int]bool)
    sum := 0

    for _, num := range nums {
        sum += num
        if sum == 0 || seen[sum] {
            return true
        }
        seen[sum] = true
    }

    return false
}
