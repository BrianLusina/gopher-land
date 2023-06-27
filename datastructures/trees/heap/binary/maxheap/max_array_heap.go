package maxheap

import (
	"errors"
	"gopherland/datastructures/trees/heap"
	"gopherland/pkg/types"
	"gopherland/pkg/utils"
)

// MaxArrayHeap represents a max heap with an array/slice as the underlying data structure
type MaxArrayHeap[T types.Comparable] struct {
	*heap.ArrayHeap[T]
}

// NewMaxArrayHeap creates a new max heap with an array as the underlying data structure
func NewMaxArrayHeap[T types.Comparable]() *MaxArrayHeap[T] {
	return &MaxArrayHeap[T]{
		&heap.ArrayHeap[T]{
			Data: []T{},
		},
	}
}

// Insert adds data to the MaxArrayHeap
func (mah *MaxArrayHeap[T]) Insert(data T) {
	// add the value into the last node
	mah.ArrayHeap.Data = append(mah.ArrayHeap.Data, data)

	// keep track of the index of the newly inserted node
	newNodeIndex := len(mah.ArrayHeap.Data) - 1

	// the following executes the "trickle up" algorithm. If the new node is not in the root position and it's greater
	// than its parent node
	for newNodeIndex > 0 && mah.ArrayHeap.Data[newNodeIndex] > mah.ArrayHeap.Data[heap.GetParentIndex(newNodeIndex)] {
		// swap the new node with the parent node
		mah.ArrayHeap.Data[heap.GetParentIndex(newNodeIndex)], mah.ArrayHeap.Data[newNodeIndex] = mah.ArrayHeap.Data[newNodeIndex], mah.ArrayHeap.Data[heap.GetParentIndex(newNodeIndex)]

		// update the index of the new node
		newNodeIndex = heap.GetParentIndex(newNodeIndex)
	}
}

// Delete removes the maximum element from the heap
func (mah *MaxArrayHeap[T]) Delete() (T, error) {
	if len(mah.Data) == 0 {
		return utils.Zero[T](), errors.New("heap is empty")
	}

	// we only ever delete the root node from a heap, so we pop the last node from the array and make it the root node
	rootNode := mah.Data[0]

	// here, we swap the first node and the last node
	mah.swap(0, len(mah.Data)-1)

	// resize to remove the root node
	mah.Data = mah.Data[:len(mah.Data)-1]

	// track the current index of the "trickle node". This is the node that will be moved into the correct position
	trickleNodeIndex := 0

	// the following loop executes the "trickle down" algorithm: We run the loop as long as the trickle node has a
	// child that is greater than it.
	for mah.hasGreaterChild(trickleNodeIndex) {
		largerChildIndex := mah.calculateLargerChildIndex(trickleNodeIndex)

		// swap the trickle node with its larger child
		mah.swap(trickleNodeIndex, largerChildIndex)

		trickleNodeIndex = largerChildIndex
	}

	return rootNode, nil
}

// hasGreaterChild checks whether the node at index has left and right children and if either of those children are greater than the node at the index.
func (mah *MaxArrayHeap[T]) hasGreaterChild(index int) bool {
	leftChildIndex := heap.GetLeftChildIndex(index)
	rightChildIndex := heap.GetRightChildIndex(index)

	leftChildExists := leftChildIndex < len(mah.Data)
	rightChildExists := rightChildIndex < len(mah.Data)

	if leftChildExists && rightChildExists {
		leftChild := mah.Data[leftChildIndex]
		rightChild := mah.Data[rightChildIndex]

		return leftChild > mah.Data[index] || rightChild > mah.Data[index]
	} else if leftChildExists && !rightChildExists {
		leftChild := mah.Data[leftChildIndex]
		return leftChild > mah.Data[index]
	} else if rightChildExists && !leftChildExists {
		rightChild := mah.Data[rightChildIndex]
		return rightChild > mah.Data[index]
	} else {
		return false
	}
}

// CalculateLargerChildIndex calculates the index of the larger child of the index of the node
func (mah *MaxArrayHeap[T]) calculateLargerChildIndex(index int) int {

	// if there is no right child
	if mah.Data[heap.GetRightChildIndex(index)] == utils.GetZeroValue[T]() {
		// return the left child index
		return heap.GetLeftChildIndex(index)
	}

	// if right child value is greater than left child value
	if mah.Data[heap.GetRightChildIndex(index)] > mah.Data[heap.GetLeftChildIndex(index)] {
		// return the right child index
		return heap.GetRightChildIndex(index)
	} else {
		return heap.GetLeftChildIndex(index)
	}
}

// swap swaps 2 elements' positions. Element at index i is moved to element at index j and vice versa
func (mah *MaxArrayHeap[T]) swap(i, j int) {
	mah.Data[i], mah.Data[j] = mah.Data[j], mah.Data[i]
}
