package disjointset

import (
	"gopherland/datastructures/graphs/vertex"
	"gopherland/pkg/types"
)

type Node[T types.Comparable] struct {
	vertex.Vertex[T]
	parent *Node[T]
	rank   int
}

// NewNode returns a new node in a disjoint set
func NewNode[T types.Comparable](data T) *Node[T] {
	node := vertex.NewVertex(data)
	n := &Node[T]{
		Vertex: *node,
		rank:   0,
	}
	n.parent = n
	return n
}

func (n *Node[T]) GetParent() *Node[T] {
	return n.parent
}

func (n *Node[T]) GetRank() int {
	return n.rank
}
