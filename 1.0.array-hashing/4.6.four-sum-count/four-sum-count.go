package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	data := make(map[int]int, len(nums1)*len(nums2))
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			data[num1+num2]++
		}
	}
	var ans int
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			ans += data[-num3-num4]
		}
	}
	return ans
}

func fourSumCount2(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := map[int]int{}

	for _, a := range nums1 {
		for _, b := range nums2 {
			target := a + b
			if _, ok := m[target]; !ok {
				m[target] = 0
			}
			m[target]++
		}
	}

	ans := 0
	for _, c := range nums3 {
		for _, d := range nums4 {
			target := -(c + d)
			if _, ok := m[target]; ok {
				ans += m[target]
			}
		}
	}
	return ans
}