package graphs

import (
	"gopherland/datastructures/graphs/vertex"
	"gopherland/pkg/types"
)

type Graph[T types.Comparable] interface {
	// AddNode adds a new node to the graph
	AddVertex(nodeOne vertex.Vertex[T]) error

	// GetNode returns a node from the graph
	GetVertex(data any) *vertex.Vertex[T]

	// Size returns the number of nodes in the graph
	Size() int
}
