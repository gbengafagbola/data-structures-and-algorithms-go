package main


func ArrayOfProducts(array []int) []int {
	n := len(array)
	result := make([]int, n)

	// First pass: prefix products
	prefix := 1
	for i := 0; i < n; i++ {
		result[i] = prefix
		prefix *= array[i]
	}

	// Second pass: suffix products
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= array[i]
	}

	return result
}

// easier to understand 
func ArrayOfProducts2(array []int) []int {
	n := len(array)
	result := make([]int, n)
	leftToRightProducts := make([]int, n)
	rightToLeftProducts := make([]int, n)

	prefix := 1
	for i := 0; i < n; i++ {
		leftToRightProducts[i] = prefix
		prefix *= array[i]
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- { 
		rightToLeftProducts[i] = suffix
		suffix *= array[i]
	}

	for i := 0; i < n; i++ { 
		result[i] = leftToRightProducts[i] * rightToLeftProducts[i]
	}

	return result
 }

func ArrayOfProducts3(array []int) []int {
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

func ArrayOfProducts4(array []int) []int {
	n := len(array)
	products := make([]int, 1, n)

	for i := range array {
		runningProduct := 1

		for j := range array {
			if i != j {
				runningProduct *= array[j]
			}
		}
	products[i] = runningProduct
	}

	return products
}