package main

func interval(intervals [][]int) [][]int {
	result := [][]int{}

	for i := 0; i < len(intervals) -1; i++ {
		nextInterval := i+1
		if i[1] >= nextInterval[0] {
			[]int{i[0], nextInterval[1]}
		}
	}
}