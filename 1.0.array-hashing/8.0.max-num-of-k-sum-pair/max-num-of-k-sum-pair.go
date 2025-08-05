package main
import "sort"

func maxOperations(nums []int, k int) int {
    sort.Ints(nums)
    i, j := 0, len(nums)-1

    count := 0
    
    for i < j {
        sum := nums[i] + nums[j] 
        if sum == k {
            count++
            i++
            j--
        } else if sum < k {
            i++
        } else {
            j--
        }
    }

    return count
}