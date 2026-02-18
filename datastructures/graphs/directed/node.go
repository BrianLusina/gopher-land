package directed

import (
	"gopherland/datastructures/graphs/vertex"
	"gopherland/pkg/types"
)

type Node[T types.Comparable] struct {
	vertex.Vertex[T]

	// InDegree is the number of edges pointing to the node
	InDegree int

	// OutDegree is the number of edges pointing away from the node
	OutDegree int
}

// NewNode creates a new node with the given data
func NewNode[T types.Comparable](data T) *Node[T] {
	node := vertex.NewVertex(data)
	return &Node[T]{
		Vertex:    *node,
		InDegree:  0,
		OutDegree: 0,
	}
}
