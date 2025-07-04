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

func BestSeat2(seats []int) int {
	bestSeat := -1
	maxSpace := 0

	left := 0 
	for left < len(seats) {
		right := left +1 
		for right < len(seats) && seats[right] == 0 {
			right += 1
		}
		availableSpace := right - left - 1
		if availableSpace > maxSpace{
			bestSeat = (left + right) / 2
			maxSpace = availableSpace
		}
		left = right
	}
	return bestSeat
}


func main() {
    seats := []int{1, 0, 0, 0, 1, 0, 1}
    fmt.Println("Best seat index:", BestSeat(seats))
}
