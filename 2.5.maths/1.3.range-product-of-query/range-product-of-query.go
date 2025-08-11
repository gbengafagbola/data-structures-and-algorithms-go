package main

func productQueries(n int, queries [][]int) []int {
	const mod = 1000000007
	var bins []int
	rep := 1
	for n > 0 {
		if n%2 == 1 {
			bins = append(bins, rep)
		}
		n /= 2
		rep *= 2
	}

	ans := make([]int, 0, len(queries))
	for _, query := range queries {
		cur := 1
		for i := query[0]; i <= query[1]; i++ {
			cur = (cur * bins[i]) % mod
		}
		ans = append(ans, cur)
	}
	return ans
}