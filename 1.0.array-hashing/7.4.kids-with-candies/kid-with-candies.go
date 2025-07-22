package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
    result := make([]bool, len(candies))

    // Step 1: Find max candies
    max := candies[0]
    for _, c := range candies {
        if c > max {
            max = c
        }
    }

    // Step 2: Check each kid
    for i, c := range candies {
        result[i] = c+extraCandies >= max
    }

    return result
}
