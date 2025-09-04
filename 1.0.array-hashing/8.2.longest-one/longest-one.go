package main

import "fmt"

func longestOnes(nums []int, k int) int {
    left := 0
    zeroCount := 0
    maxLen := 0

    for right := 0; right < len(nums); right++ {
        if nums[right] == 0 {
            zeroCount++
        }

        // shrink window if zeros exceed k
        for zeroCount > k {
            if nums[left] == 0 {
                zeroCount--
            }
            left++
        }

        // update max length
        maxLen = max(maxLen, right-left+1)
    }

    return maxLen
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
    k := 2
    fmt.Println(longestOnes(nums, k)) // Output: 6
}
