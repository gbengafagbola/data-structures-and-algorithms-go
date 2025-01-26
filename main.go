package main

import "fmt"

type Queue struct {
	items chan int
}

func (q *Queue) Enqueue(item int) {
	q.items <- item
}

func (q *Queue) Dequeue() int {
		if len(q.items) == 0 {
			return -1
		}
		return <- q.items
}

func main() {
	q := Queue{
		items: make(chan int, 100),
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
	fmt.Println(q.Dequeue())
 
}
