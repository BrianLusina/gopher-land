package binary

import "gopherland/pkg/types"

type Heap[T types.Comparable] []T

// Len returns the length of the heap
func (h Heap[T]) Len() int {
	return len(h)
}

// Less returns true if element at position i is less than element at position j
func (h Heap[T]) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap swaps the elements at position i & j
func (h Heap[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds an element to the heap
// Using a pointer receiver because this modifies the slice's length
func (h *Heap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

// Pop removes the last element from the heap
// Using a pointer receiver because this modifies the slice's length
func (h *Heap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Peek checks the top element without removing it
func (h *Heap[T]) Peek() T {
	p := *h
	top := p[0]
	return top
}
