package main

import (
    "fmt"
    "sort"
)

func ZeroSumSubarray(nums []int) bool {
    l := 0
    r := len(nums) - 1

    sort.Ints(nums)

    for l < r {
        sum := nums[l] + nums[r]

        if sum == 0 {
            return true
        } else if sum < 0 {
            if r-1 > l {
                newSum := sum + nums[r-1]
                if newSum == 0 {
                    return true
                }
            }
            l++
        } else {
            if l+1 < r {
                newSum := sum + nums[l+1]
                if newSum == 0 {
                    return true
                }
            }
            r--
        }
    }

    return false
}

func main() {
    fmt.Println(ZeroSumSubarray([]int{-5, -5, -2, 2, 3})) // true
}
