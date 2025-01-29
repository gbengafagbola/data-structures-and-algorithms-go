package main

import "fmt"

type Queue struct {
	items chan int
}

func (q *Queue) Enqueue(item int) {
	q.items <- item
}

// Channels in Go have FIFO (First In, First Out) behavior, meaning that the first item sent to the channel will be the first one received.
func (q *Queue) Dequeue() int {
		return <- q.items
}

func main() {
	q := Queue{
		items: make(chan int, 16),
	}

	q.Enqueue(1)
	q.Enqueue(3)
	q.Enqueue(5)
	q.Enqueue(7)

	// fmt.Println("deque:")
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
 
}
