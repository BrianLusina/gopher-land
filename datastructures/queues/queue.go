package queues

import "errors"

var (
	ErrorEmptyQueue = errors.New("queue is empty")
	ErrorFullQueue  = errors.New("queue is full")
)

type Queue interface {
	Enqueue(x interface{})
	Dequeue() (interface{}, error)
	Peek() (interface{}, error)
}
