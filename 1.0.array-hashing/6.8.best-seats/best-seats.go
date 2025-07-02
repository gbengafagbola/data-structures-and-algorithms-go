package main

import "fmt"

func BestSeat(seats []int) int {
    result := -1
    diff := -1

    for i := 1; i < len(seats)-1; {
        j := i
        if seats[i] == 0 && seats[j] == 0 {
            for j < len(seats) && seats[j] == 0 {
                j++
            }
            mid := (i + j - 1) / 2
            difference := j - i
            if difference > diff {
                diff = difference
                result = mid
            }
            i = j
        } else {
            i++
        }
    }

    return result
}

func main() {
    seats := []int{1, 0, 0, 0, 1, 0, 1}
    fmt.Println("Best seat index:", BestSeat(seats))
}
