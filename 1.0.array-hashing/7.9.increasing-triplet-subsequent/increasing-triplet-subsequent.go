package main

func increasingTriplet(nums []int) bool {
    first := int(^uint(0) >> 1)  // MaxInt
    second := int(^uint(0) >> 1) // MaxInt

    for _, n := range nums {
        if n <= first {
            first = n
        } else if n <= second {
            second = n
        } else {
            return true
        }
    }

    return false
}
