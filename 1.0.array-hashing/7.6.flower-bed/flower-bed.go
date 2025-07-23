package main


func canPlaceFlowers(flowerbed []int, n int) bool {
    numberOfConsecutiveThreeZero := 0

    for i := 0; i < len(flowerbed); i++ {
        if flowerbed[i] == 0 &&
            (i == 0 || flowerbed[i-1] == 0) &&
            (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {

            flowerbed[i] = 1 // simulate planting
            numberOfConsecutiveThreeZero++
        }
    }
    return numberOfConsecutiveThreeZero >= n
}


func canPlaceFlowersInitial(flowerbed []int, n int) bool {
    numberOfConsecutiveThreeZero := 0
    threeZeroes := 0

    for i := 0; i < len(flowerbed); i++ {
        // Count zeros
        if flowerbed[i] == 0 {
            threeZeroes++
        } else {
            threeZeroes = 0 // reset if you see a 1
        }

        // Check if we have 3 zeros in a row (or at edge)
        if threeZeroes == 3 {
            numberOfConsecutiveThreeZero++
            threeZeroes = 1 // reset to 1 to allow overlapping 000 segments like 0000
        }
    }

    return numberOfConsecutiveThreeZero >= n
}
