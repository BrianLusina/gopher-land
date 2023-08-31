package edges

import (
	"gopherland/datastructures/graphs"
	"gopherland/pkg/types"
)

// Edge represents an edge in a graph
type Edge[T types.Comparable] interface {
	// IsSelfEdge is used to check if an edge has a vertex pointing to itself
	IsSelfEdge() bool

	// GetNodes retrieves all the nodes in this edge
	GetNodes() []*graphs.Vertex[T]

	// IsWeighted is used to check if an edge has a weight/cost
	IsWeighted() bool

	// GetWeight retrieve the weight of the edge
	GetWeight() int
}
