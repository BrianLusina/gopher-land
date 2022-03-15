package fifo

import (
	"gopherland/datastructures/queues"
)

type FifoQueue struct {
	items []interface{}
}

func NewQueue() queues.Queue {
	// new built in function can be used or struct literal
	// return &FifoQueue{}
	return new(FifoQueue)
}

// Enqueue adds an item to the back of the queue
func (q *FifoQueue) Enqueue(x interface{}) {
	q.items = append(q.items, x)
}

// Dequeue removes an item from the front of the queue
func (q *FifoQueue) Dequeue() interface{} {
	if q.Empty() {
		panic("Queue is empty")
	}
	value := q.items[0]
	q.items = q.items[1:]
	return value
}

// Peek returns the front element of the queue without removing it
func (q *FifoQueue) Peek() interface{} {
	if q.Empty() {
		panic("Queue is empty")
	}
	return q.items[0]
}

// Empty returns true if the queue is empty
func (q *FifoQueue) Empty() bool {
	return len(q.items) == 0
}

// Items returns a slice of the items in the queue
func (q *FifoQueue) Items() []interface{} {
	return q.items
}
