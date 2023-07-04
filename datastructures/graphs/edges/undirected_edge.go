package edges

import (
	"gopherland/datastructures/graphs"
	"gopherland/pkg/types"
)

// UndirectedEdge represents an undirected edge in a graph. This means that one can move from one node to another & vice versa along this edge
type Undirected[T types.Comparable] struct {
	// NodeOne is 1 node in the edge
	NodeOne *graphs.Vertex[T]

	// NodeTwo is the other node in the edge
	NodeTwo *graphs.Vertex[T]

	// Weight is the weight of the edge, is 0, this is considered an unweighted undirected edge in a graph
	Weight int
}

// NewUndirectedEdge creates a new undirected edge between two nodes with a weight
func NewUndirectedEdge[T types.Comparable](nodeOne, nodeTwo *graphs.Vertex[T], weight int) Edge[T] {
	return &Undirected[T]{
		NodeOne: nodeOne,
		NodeTwo: nodeTwo,
		Weight:  weight,
	}
}

// IsSelfEdge returns true if an edge's 2 nodes are the same
func (e *Undirected[T]) IsSelfEdge() bool {
	return e.NodeOne == e.NodeTwo
}

// GetNodes returns the nodes connected by the edge
// this includes the nodes in the edges if any
func (e *Undirected[T]) GetNodes() []*graphs.Vertex[T] {
	return []*graphs.Vertex[T]{e.NodeOne, e.NodeTwo}
}

// IsWeighted returns true if the edge has a weight
func (e *Undirected[T]) IsWeighted() bool {
	return e.Weight != 0
}

// IsWeighted returns true if the edge has a weight
func (e *Undirected[T]) GetWeight() int {
	return e.Weight
}
