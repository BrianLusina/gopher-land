package graphs

import "gopherland/pkg/types"

type Graph[T types.Comparable] interface {
	// AddNode adds a new node to the graph
	AddVertex(nodeOne Vertex[T]) error

	// GetNode returns a node from the graph
	GetVertex(data any) *Vertex[T]

	// Size returns the number of nodes in the graph
	Size() int
}
