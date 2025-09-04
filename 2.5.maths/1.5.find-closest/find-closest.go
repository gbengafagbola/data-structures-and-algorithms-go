package main
import "math"

func findClosest(x int, y int, z int) int {
    rate1 := int(math.Max(float64(z - x), float64(x - z)))
    rate2 := int(math.Max(float64(z - y), float64(y - z)))

    if rate1 < rate2 {
        return 1
    } else if rate1 > rate2 {
        return 2
    } else {
    return 0
    }

}