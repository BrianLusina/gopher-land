package maxheap

import (
	"gopherland/datastructures/trees/heap"
	"gopherland/pkg/types"
)

// MaxArrayHeap represents a max heap with an array/slice as the underlying data structure
type MaxArrayHeap[T types.Comparable] struct {
	heap.ArrayHeap[T]
}

// NewMaxArrayHeap creates a new max heap with an array as the underlying data structure
func NewMaxArrayHeap[T types.Comparable]() *MaxArrayHeap[T] {
	return &MaxArrayHeap[T]{
		heap.ArrayHeap[T]{
			Data: []T{},
		},
	}
}
