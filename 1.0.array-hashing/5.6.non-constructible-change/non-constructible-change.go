package main
import "sort"

func NonConstructibleChange(coins []int) int {
    sort.Ints(coins)

    if 1 != coins[0] {
        return 1
    } 
    
    result := []int{}
    m := make(map[int]bool)

    for index, coin := range coins {
        if index+1 == coins[index] || addUpTo(index+1) {
            m[index+1]=true
        } else {
            return index+1
        }
        
    }

    return 0
}


func addUpTo(sum int, coins []int) bool {
    result := 0
    for _, coin := range coins {
        result += coin
        if result == sum {
            return true
        }
    }
    return false
}