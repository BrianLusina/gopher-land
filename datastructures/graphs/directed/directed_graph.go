package directed

import "gopherland/datastructures/graphs"

// DirectedGraph is a directed graph
type DirectedGraph struct {
	root *Node

	// nodes is a map of the nodes in the graph
	nodes map[any]*Node

	// count is the number of nodes in the graph
	count int
}

func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{
		nodes: make(map[any]*Node),
	}
}

func (g *DirectedGraph) AddNode(node *Node) error {
	if g.nodes == nil {
		g.nodes = make(map[any]*Node)
	}

	if g.nodeInGraph(node) {
		return graphs.ErrNodeExists
	} else {
		g.nodes[node.Data] = node
		g.count++
		return nil
	}
}

func (g *DirectedGraph) GetNode(data any) *Node {
	for _, node := range g.nodes {
		if node.Data == data {
			return node
		}
	}
	return nil
}

func (g *DirectedGraph) Size() int {
	return g.count
}

func (g *DirectedGraph) nodeInGraph(node *Node) bool {
	_, ok := g.nodes[node.Data]
	return ok
}
