package undirected

import "gopherland/datastructures/graphs"

// UnDirectedGraph
type UnDirectedGraph struct {
	root *graphs.Node
	// nodes is a slice of the nodes in the graph
	nodes map[any]*graphs.Node
	// edges is a map of the edges in the graph, where the key is the data of the node and
	// the value is the slice of the node's neighbours
	edges map[any][]*graphs.Node
	// count is the number of nodes in the graph
	count int
}

func NewGraph() *UnDirectedGraph {
	return &UnDirectedGraph{
		nodes: make(map[any]*graphs.Node),
	}
}

func (g *UnDirectedGraph) AddNode(node graphs.Node) error {
	if g.edges == nil {
		g.edges = make(map[any][]*graphs.Node)
	}

	if _, ok := g.nodes[node.Data]; ok {
		return graphs.ErrNodeExists
	} else {
		if node.Neighbors != nil {
			for _, neighbour := range node.Neighbors {
				g.edges[node.Data] = append(g.edges[node.Data], neighbour)
			}
		} else {
			g.edges[node.Data] = []*graphs.Node{}
		}

		g.nodes[node.Data] = &node
		g.count++
		return nil
	}
}

func (g *UnDirectedGraph) GetNode(data any) *graphs.Node {
	for _, node := range g.nodes {
		if node.Data == data {
			return node
		}
	}
	return nil
}

func (g *UnDirectedGraph) GetEdges() map[any][]*graphs.Node {
	return g.edges
}

func (g *UnDirectedGraph) Size() int {
	return g.count
}
