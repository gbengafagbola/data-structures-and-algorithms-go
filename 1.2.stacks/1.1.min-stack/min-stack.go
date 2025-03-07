package main

import "math"

type MinStack struct {
	stack [][]int 
}


func Constructor() MinStack {
	return MinStack{
		stack: [][]int{},
	}
}

func (ms *MinStack) Push(x int)  {
	currentMinimum := x 

	if len(ms.stack) > 0 {
		lastElement := ms.stack[len(ms.stack) - 1]
		currentMinimum = int(math.Min(float64(currentMinimum), float64(lastElement[1])))
	}
	ms.stack = append(ms.stack, []int{x, currentMinimum})
}

func (ms *MinStack) Pop()  {
	ms.stack = ms.stack[:len(ms.stack) - 1]
}

func (ms *MinStack) Top() int {
	if len(ms.stack) == 0 {
		return -1
	}
	lastElement := ms.stack[len(ms.stack) - 1]
	return lastElement[0]
}


func (ms *MinStack) GetMin() int {
	if len(ms.stack) == 0 {
		return -1
	}
	lastElement := ms.stack[len(ms.stack) - 1]
	return lastElement[1]
}