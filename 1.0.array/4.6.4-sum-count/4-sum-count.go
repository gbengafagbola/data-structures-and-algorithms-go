package main

func fourSumCount(A []int, B []int, C []int, D []int) int {
	m := map[int]int{}

	for _, a := range A {
		for _, b := range B {
			target := a + b
			if _, ok := m[target]; !ok {
				m[target] = 0
			}
			m[target]++
		}
	}

	ans := 0
	for _, c := range C {
		for _, d := range D {
			target := -(c + d)
			if _, ok := m[target]; ok {
				ans += m[target]
			}
		}
	}
	return ans
}
