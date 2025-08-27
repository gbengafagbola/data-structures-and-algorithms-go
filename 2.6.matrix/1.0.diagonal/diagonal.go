package main

func areaOfMaxDiagonal(dimensions [][]int) int {
    var maxDiagSq int64 = -1
    var bestArea  int64 = 0

    for _, d := range dimensions {
        a := int64(d[0])
        b := int64(d[1])

        diagSq := a*a + b*b  // diagonal squared
        area   := a*b        // rectangle area

        if diagSq > maxDiagSq || (diagSq == maxDiagSq && area > bestArea) {
            maxDiagSq = diagSq
            bestArea  = area
        }
    }
    return int(bestArea)
}
