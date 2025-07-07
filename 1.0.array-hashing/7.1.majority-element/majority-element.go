package main
import "sort"

func MajorityElement(array []int) int {
    candidate := array[0]
    count := 1

    // Find candidate using Boyer-Moore Voting Algorithm
    for i := 1; i < len(array); i++ {
        if count == 0 {
            candidate = array[i]
        }
        if array[i] == candidate {
            count++
        } else {
            count--
        }
    }

    // Verify candidate
    count = 0
    for _, num := range array {
        if num == candidate {
            count++
        }
    }
    if count > len(array)/2 {
        return candidate
    }
    return -1
}

func MajorityElement2(array []int) int {
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
