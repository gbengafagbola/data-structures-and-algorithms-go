package main
import "sort"


func maxOperations(nums []int, k int) int {
    count := 0
    freq := make(map[int]int)

    for _, num := range nums {
        complement := k - num
        if freq[complement] > 0 {
            count++
            freq[complement]-- // use the complement
        } else {
            freq[num]++ // store num for future pairing
        }
    }

    return count
}

func maxOperationsInitial(nums []int, k int) int {
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