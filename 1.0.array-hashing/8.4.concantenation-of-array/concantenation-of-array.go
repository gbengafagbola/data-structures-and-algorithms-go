func getConcatenation(nums []int) []int {
    length := len(nums)
    ans := make([]int, 2*length)

    for i := 0; i < length; i++ {
        ans[i] = nums[i]
        ans[i+length] = nums[i]
    }

    return ans
}


// func getConcatenation(nums []int) []int {
//     ans := []int{}

//     for i := 0; i < 2; i++ {
//         ans = append(ans, nums...)
//     }

//     return ans
// }

// func getConcatenation(nums []int) []int {
    
//     length := len(nums)

//     ans := []int{}
//     // ans := make([]int, 2*length)

//     for i := 0; i < length; i++ {
//         ans = append(ans, nums[i])
//     }

//     for i := 0; i < length; i++ {
//         ans = append(ans, nums[i])
//     }

//     return ans
// }