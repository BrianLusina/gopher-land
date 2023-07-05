package edges

import (
	"gopherland/datastructures/graphs"
	"gopherland/pkg/types"
)

// Hyper represents a hyper edge in a graph which connects many vertices
type Hyper[T types.Comparable] struct {
	// nodes is a slice of all the nodes this edge connects
	nodes []*graphs.Vertex[T]
}

// NewHyperEdge creates a new directed edge between two nodes with a weight
func NewHyperEdge[T types.Comparable](nodes ...*graphs.Vertex[T]) Edge[T] {
	n := []*graphs.Vertex[T]{}
	n = append(n, nodes...)
	return &Hyper[T]{
		nodes: n,
	}
}

// IsSelfEdge returns true if an edge's 2 nodes are the same
func (e *Hyper[T]) IsSelfEdge() bool {
	return false
}

// GetNodes returns the nodes connected by the edge
// this includes the nodes in the edges if any
func (e *Hyper[T]) GetNodes() []*graphs.Vertex[T] {
	return e.nodes
}

// IsWeighted returns true if the edge has a weight
func (e *Hyper[T]) IsWeighted() bool {
	return false
}

// IsWeighted returns true if the edge has a weight
func (e *Hyper[T]) GetWeight() int {
	return 0
}
