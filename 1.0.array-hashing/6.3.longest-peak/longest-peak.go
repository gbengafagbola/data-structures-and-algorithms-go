package main

func LongestPeak(array []int) int {
	maxPeakLength := 0
	n := len(array)
	i := 1 // Start from second element

	for i < n-1 {
		// Check if current element is a peak tip
		if array[i-1] < array[i] && array[i] > array[i+1] {
			left := i - 2
			for left >= 0 && array[left] < array[left+1] {
				left--
			}

			right := i + 2
			for right < n && array[right] < array[right-1] {
				right++
			}

			currentPeakLength := right - left - 1
			if currentPeakLength > maxPeakLength {
				maxPeakLength = currentPeakLength
			}
			i = right // Skip to the end of this peak
		} else {
			i++
		}
	}
	return maxPeakLength
}

