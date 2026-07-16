package undirected

import (
	"gopherland/datastructures/graphs"
	"gopherland/pkg/types"
)

// UnDirectedGraph
type UnDirectedGraph[T types.Comparable] struct {
	root *Node[T]
	// nodes is a slice of the nodes in the graph
	nodes map[any]*Node[T]
	// edges is a map of the edges in the graph, where the key is the data of the node and
	// the value is the slice of the node's neighbours
	edges map[any][]*Node[T]
	// count is the number of nodes in the graph
	count int
}

func NewGraph[T types.Comparable]() *UnDirectedGraph[T] {
	return &UnDirectedGraph[T]{
		nodes: make(map[any]*Node[T]),
	}
}

func (g *UnDirectedGraph[T]) AddNode(node Node[T]) error {
	if g.edges == nil {
		g.edges = make(map[any][]*Node[T])
	}

	if _, ok := g.nodes[node.data]; ok {
		return graphs.ErrNodeExists
	} else {
		if node.Neighbors != nil {
			for _, neighbour := range node.Neighbors {
				g.edges[node.data] = append(g.edges[node.data], neighbour)
			}
		} else {
			g.edges[node.data] = []*graphs.Node{}
		}

		g.nodes[node.data] = &node
		g.count++
		return nil
	}
}

func (g *UnDirectedGraph[T]) GetNode(data any) *Node[T] {
	for _, node := range g.nodes {
		if node.data == data {
			return node
		}
	}
	return nil
}

func (g *UnDirectedGraph[T]) GetEdges() map[any][]*Node[T] {
	return g.edges
}

func (g *UnDirectedGraph[T]) Size() int {
	return g.count
}
