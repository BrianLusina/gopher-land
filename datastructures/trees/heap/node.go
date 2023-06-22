package heap

import (
	"gopherland/datastructures/trees/binary"
	"gopherland/pkg/types"
)

// HeapNode represents a node in a Heap
type HeapNode[T types.Comparable] struct {
	binary.BinaryTreeNode[T]
}
