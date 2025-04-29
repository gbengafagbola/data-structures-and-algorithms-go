package main


func IsMonotonic(array []int) bool {
    if len(array) <= 2 {
        return true
    }

    direction := array[1] - array[0]

    for i := 2; i < len(array); i++ {
        if direction == 0 {
            direction = array[i] - array[i-1]
            continue
        }

        if breaksDirection(direction, array[i-1], array[i]) {
            return false
        }
    }

    return true
}


func IsMonotonic2(array []int) bool {
    if len(array) <= 2 {
        return true
    }

    direction := array[1] - array[0]

    for i := 2; i < len(array); i++ {
        if direction == 0 {
            direction = array[i] - array[i-1]
            continue
        }

        if breaksDirection(direction, array[i-1], array[i]) {
            return false
        }
    }

    return true
}

func breaksDirection(direction int, prev int, curr int) bool {
    diff := curr - prev
    return direction > 0 && diff < 0 || direction < 0 && diff > 0
}


func IsMonotonic3(array []int) bool {
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
