func shuffle(nums []int, n int) []int {
    result := make([]int, 2*n)

    for i := 0; i < n; i++ {
        result[2*i] = nums[i]
        result[2*i+1] = nums[i+n]
    }

    return result
}


func shuffle1(nums []int, n int) []int {
    result := make([]int, 0, len(nums))

    for i := 0; i < n; i++ {
        result = append(result, nums[i])
        result = append(result, nums[i+n])
    }

    return result
}