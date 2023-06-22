package heap

import "gopherland/pkg/types"

// Heap is an interface that contains the method set available for heaps
type Heap[T types.Comparable] interface {
	// Insert adds a new data element to the heap
	Insert(data T)

	// Deletes a node from the heap and returns its value
	Delete() T
}
