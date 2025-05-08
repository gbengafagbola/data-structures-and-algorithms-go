package main

import "math"

func LongestPeak(array []int) int {
	i := 1
	n := len(array)
	longestPeakLength := 0

	for i < n-1 {
		isPeak := array[i] > array[i-1] && array[i] > array[i+1]
		if !isPeak {
			i++
			continue
		}

		leftIdx := i - 2
		for leftIdx >= 0 && array[leftIdx] < array[leftIdx+1] {
			leftIdx--
		}

		rightIdx := i + 2
		for rightIdx < n && array[rightIdx] < array[rightIdx-1] {
			rightIdx++
		}

		currentPeakLength := rightIdx - leftIdx - 1
		longestPeakLength = int(math.Max(float64(currentPeakLength), float64(longestPeakLength)))

		// Move i to the end of the current peak to avoid re-processing
		i = rightIdx
	}

	return longestPeakLength
}


func LongestPeak2(array []int) int {

	i := 1
	n := len(array)
	peaks := []int{}
	longestPeakLength := 0

	for i < n-1 {
		if array[i] > array[i-1] && array[i] > array[i+1] {
			peaks = append(peaks, i)
		}
		i++
	}

	for peak := range peaks {
		i := 1
		tempPeak := 0
		if array[peak] > array[peak-i] && array[peak] > array[peak+i] {
			tempPeak++
		}
		i++
		if tempPeak < longestPeakLength {
			longestPeakLength = tempPeak
		}
	}
	return longestPeakLength

}

func LongestPeak3(array []int) int {
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
