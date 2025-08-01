package main
import "fmt"

func compress(chars []byte) int {
    result := []byte{}
    count := 1

    for i := 0; i < len(chars); i++ {
        if i+1 < len(chars) && chars[i] == chars[i+1] {
            count++
        } else if count == 1 {
            result = append(result, chars[i])
            count = 1
        } else if count > 1 && count < 10 {
            result = append(result, chars[i])
            result = append(result, byte('0'+count))
            count = 1
        } else if count >= 10 {
            result = append(result, chars[i])
            countStr := []byte(fmt.Sprintf("%d", count))
            for j := 0; j < len(countStr); j++ {
                result = append(result, countStr[j])
            }
            count = 1
        }
    }

    // Optional: copy result back into original array if needed
    for i := 0; i < len(result); i++ {
        chars[i] = result[i]
    }

    return len(result)
}
