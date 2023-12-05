package disjointset

import "gopherland/pkg/types"

type DisjointSet[T types.Comparable] struct {
	nodes []*Node[T]
	size  int
}

func NewDisjointSet[T types.Comparable](node ...Node[T]) *DisjointSet[T] {
	nodes := make([]*Node[T], len(node))
	size := len(node)

	for i, n := range node {
		nodes[i] = &n
	}

	return &DisjointSet[T]{
		size:  size,
		nodes: nodes,
	}
}

// Union takes two nodes and makes them part of the same set
// set with bigger rank should be parent, so that the
// disjoint set tree will be more flat.
func (d *DisjointSet[T]) Union(nodeOne, nodeTwo *Node[T]) {
	rootX := d.Find(nodeOne)
	rootY := d.Find(nodeTwo)

	if rootX != rootY {
		if rootX.rank > rootY.rank {
			rootY.parent = rootX
		} else if rootX.rank < rootY.rank {
			rootX.parent = rootY
		} else {
			rootX.parent = rootY
			rootY.rank++
		}
	}
}

// Find returns the root node of the node
func (d *DisjointSet[T]) Find(node *Node[T]) *Node[T] {
	if node != node.parent {
		node.parent = d.Find(node.parent)
	}
	return node.parent
}

func (d *DisjointSet[T]) Connected(nodeOne, nodeTwo *Node[T]) bool {
	return d.Find(nodeOne) == d.Find(nodeTwo)
}

func (d *DisjointSet[T]) Size() int {
	return d.size
}
