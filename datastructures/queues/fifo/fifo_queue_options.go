package fifo

type FifoQueueOptions[T any] func(*FifoQueue[T])

// WithMaxSize sets the maximum size of the queue
func WithMaxSize[T any](maxSize int) FifoQueueOptions[T] {
	return func(q *FifoQueue[T]) {
		q.maxSize = maxSize
	}
}
