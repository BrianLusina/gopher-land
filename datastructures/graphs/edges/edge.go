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

// // NewEdge creates a new edge between two nodes with a weight
// func NewEdge(nodeOne, nodeTwo *Vertex, weight int) *Edge {
// 	if nodeOne.Neighbors == nil {
// 		nodeOne.Neighbors = map[any]*Node{}
// 		nodeOne.Neighbors[nodeTwo.Data] = nodeTwo
// 	} else {
// 		// is nodeTwo already in nodeOne's Neighbors?
// 		if _, ok := nodeOne.Neighbors[nodeTwo.Data]; !ok {
// 			nodeOne.Neighbors[nodeTwo.Data] = nodeTwo
// 		}
// 	}

// 	if nodeTwo.Neighbors == nil {
// 		nodeTwo.Neighbors = map[any]*Node{}
// 		nodeTwo.Neighbors[nodeOne.Data] = nodeOne
// 	} else {
// 		if _, ok := nodeTwo.Neighbors[nodeOne.Data]; !ok {
// 			nodeTwo.Neighbors[nodeOne.Data] = nodeOne
// 		}
// 	}

// 	// increment the degree of both nodes
// 	nodeOne.Degree++
// 	nodeTwo.Degree++

// 	return &Edge{
// 		NodeOne: nodeOne,
// 		NodeTwo: nodeTwo,
// 		Weight:  weight,
// 		edges:   []*Edge{},
// 	}
// }

// // AddEdge adds an edges to an edge
// func (e *Edge) AddEdge(edge ...*Edge) {
// 	e.edges = append(e.edges, edge...)
// }

// // GetNodes returns the nodes connected by the edge
// // this includes the nodes in the edges if any
// func (e *Edge) GetNodes() []*Node {
// 	if len(e.edges) == 0 {
// 		return []*Node{e.NodeOne, e.NodeTwo}
// 	}
// 	nodes := []*Node{e.NodeOne, e.NodeTwo}

// 	for _, edge := range e.edges {
// 		nodes = append(nodes, edge.GetNodes()...)
// 	}

// 	return nodes
// }
