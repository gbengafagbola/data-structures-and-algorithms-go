package main

func ArrayOfProducts(array []int) []int {
	result := []int{}

	for i := 0; i < len(array); i++ {
		currentProduct := 1
		x := []int{} // reset x on each iteration

		// Append all elements after i
		if i != len(array)-1 {
			x = append(x, array[i+1:]...)
		}

		// Append all elements before i
		if i != 0 {
			x = append(x, array[:i]...)
		}

		// Multiply all values in x
		for _, y := range x {
			currentProduct *= y
		}

		result = append(result, currentProduct)
	}
	return result
}
