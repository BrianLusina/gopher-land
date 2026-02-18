package directed

import (
	"gopherland/datastructures/graphs"
	"gopherland/pkg/types"
)

// DirectedGraph is a directed graph
type DirectedGraph[T types.Comparable] struct {
	root *Node[T]

	// nodes is a map of the nodes in the graph
	nodes map[any]*Node[T]

	// count is the number of nodes in the graph
	count int
}

func NewDirectedGraph[T types.Comparable]() *DirectedGraph[T] {
	return &DirectedGraph[T]{
		nodes: make(map[any]*Node[T]),
	}
}

func (g *DirectedGraph[T]) AddNode(node *Node[T]) error {
	if g.nodes == nil {
		g.nodes = make(map[any]*Node[T])
	}

	if g.nodeInGraph(node) {
		return graphs.ErrNodeExists
	} else {
		g.nodes[node.Data()] = node
		g.count++
		return nil
	}
}

func (g *DirectedGraph[T]) GetNode(data any) *Node[T] {
	for _, node := range g.nodes {
		if node.Data() == data {
			return node
		}
	}
	return nil
}

func (g *DirectedGraph[T]) Size() int {
	return g.count
}

func (g *DirectedGraph[T]) nodeInGraph(node *Node[T]) bool {
	_, ok := g.nodes[node.Data()]
	return ok
}
