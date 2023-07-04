package edges

import (
	"gopherland/datastructures/graphs"
	"gopherland/pkg/types"
)

// Hyper represents a hyper edge in a graph which connects many vertices
type Hyper[T types.Comparable] struct {
	// nodes is a slice of all the nodes this edge connects
	nodes *graphs.Vertex[T]
}

// NewDirectedEdge creates a new directed edge between two nodes with a weight
func NewHyperEdge[T types.Comparable](nodes ...*graphs.Vertex[T]) *Directed[T] {
	return &Directed[T]{
		Source:      source,
		Destination: destination,
	}
}

// IsSelfEdge returns true if an edge's 2 nodes are the same
func (e *Directed[T]) IsSelfEdge() bool {
	return e.NodeOne == e.NodeTwo
}

// GetNodes returns the nodes connected by the edge
// this includes the nodes in the edges if any
func (e *Directed[T]) GetNodes() []*graphs.Vertex[T] {
	return []*graphs.Vertex[T]{e.NodeOne, e.NodeTwo}
}

// IsWeighted returns true if the edge has a weight
func (e *Directed[T]) IsWeighted() bool {
	return e.Weight != 0
}
