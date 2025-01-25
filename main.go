package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(val int) {
	q.items = append(q.items, val)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}
	item, items := q.items[0], q.items[1:]
	q.items = items
	return item
}

func main() {
	q := Queue{}

	q.Enqueue(1)
	q.Enqueue(3)
	q.Enqueue(5)
	q.Enqueue(7)

	// fmt.Println("deque:")
	fmt.Println(q.items)
 
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
