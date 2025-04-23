package main

func IsMonotonic(array []int) bool {
    if len(array) <= 2 {
        return true
    }

    count := 0
    equals := 0

    for i := 0; i < len(array)-1; i++ {
        j := i + 1
        if array[i] < array[j] {
            count--
        } else if array[i] > array[j] {
            count++
        } else {
            equals++
        }
    }

    if count < 0 {
        count = -count
    }

    totalComparisons := count + equals
    return totalComparisons == len(array)-1
}
