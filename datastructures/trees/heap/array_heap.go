package heap

import (
	"errors"
	"gopherland/pkg/types"
	"gopherland/pkg/utils"
)

// ArrayHeap is an array based heap
type ArrayHeap[T types.Comparable] struct {
	Data []T // underlying data structure used to maintain the heap, a slice is used in this case
}

// returns the root node value in a heap
func (a *ArrayHeap[T]) RootNode() (T, error) {
	if len(a.Data) == 0 {
		return utils.Zero[T](), errors.New("heap is empty")
	}
	return a.Data[0], nil
}

// LastNode returns the last node in a Heap
func (a *ArrayHeap[T]) LastNode() (T, error) {
	if len(a.Data) == 0 {
		return utils.Zero[T](), errors.New("heap is empty")
	}
	return a.Data[len(a.Data)-1], nil
}
