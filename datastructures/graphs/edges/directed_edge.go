package edges

import (
	"gopherland/datastructures/graphs"
	"gopherland/pkg/types"
)

// Directed represents a directed edge in a graph. This means that one can only move from one node to another & NOT vice versa along this edge
type Directed[T types.Comparable] struct {
	// Source is the source node in the edge
	Source *graphs.Vertex[T]

	// Destination is the destination node in the edge
	Destination *graphs.Vertex[T]

	// Weight is the weight of the edge, if 0, this is considered an unweighted directed edge in a graph
	Weight int
}

// NewDirectedEdge creates a new directed edge between two nodes with a weight
func NewDirectedEdge[T types.Comparable](source, destination *graphs.Vertex[T], weight int) *Directed[T] {
	return &Directed[T]{
		Source:      source,
		Destination: destination,
		Weight:      weight,
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
