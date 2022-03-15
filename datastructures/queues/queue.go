package queues

type Queue interface {
	Enqueue(x interface{})
	Dequeue() interface{}
	Peek() interface{}
	Empty() bool
	Items() []interface{}
}
