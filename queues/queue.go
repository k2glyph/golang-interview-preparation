package main

import "fmt"

// Queue stack which hold all the items
type Queue struct {
	items []int
}

// Enqueue
func (q *Queue) enQueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue
func (q *Queue) deQueue() int {
	itemToRemove := q.items[0]
	q.items = q.items[1:]
	return itemToRemove
}
func main() {
	myQueue := Queue{}
	myQueue.enQueue(10)
	myQueue.enQueue(20)
	myQueue.enQueue(30)
	fmt.Println(myQueue)
	myQueue.deQueue()
	fmt.Println(myQueue)
	myQueue.deQueue()
	fmt.Println(myQueue)
}
