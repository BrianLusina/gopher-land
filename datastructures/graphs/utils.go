package graphs

import (
	"gopherland/datastructures/graphs/vertex"
	"gopherland/pkg/types"
)

// CloneGraph clones a graph and returns a new graph
// this assumes the graph is connected and undirected
func CloneGraph[T types.Comparable](node *vertex.Vertex[T]) *vertex.Vertex[T] {
	if node == nil {
		return nil
	}

	m := map[any]*vertex.Vertex[T]{
		node.data: vertex.NewVertex(node.data),
	}

	queue := []*vertex.Vertex[T]{node}

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbour := range current.GetNeighbors() {
			if _, ok := m[neighbour.data]; !ok {
				queue = append(queue, neighbour)
				m[neighbour.data] = vertex.NewVertex(neighbour.data)
			}
			m[current.data] = m[neighbour.data]
		}
	}

	return m[node.data]
}
