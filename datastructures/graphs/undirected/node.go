package undirected

import (
	"gopherland/datastructures/graphs/vertex"
	"gopherland/pkg/types"
)

type Node[T types.Comparable] struct {
	vertex.Vertex[T]
}
