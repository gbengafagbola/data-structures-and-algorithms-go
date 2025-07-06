package main
import "sort"
func MajorityElement(array []int) int {
    n := make(map[int]int)

    for _, num := range array {
        if _, seen := n[num]; seen {
            n[num] += 1 
        } else {
        n[num] = 1
            
        }
    }

    occurence := []int{}
    for _, value := range n {
        occurence = append(occurence, value)
    }

    sort.Ints(occurence)
    v := occurence[len(occurence)-1]
    if v >= len(array)/2 {
        
     for key, value := range n {
		if value == v {
			return key
		}
	}
    }
    return -1
}
