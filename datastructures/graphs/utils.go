package graphs

import "gopherland/pkg/types"

// CloneGraph clones a graph and returns a new graph
// this assumes the graph is connected and undirected
func CloneGraph[T types.Comparable](node *Vertex[T]) *Vertex[T] {
	if node == nil {
		return nil
	}

	m := map[any]*Vertex[T]{
		node.Data: NewVertex(node.Data),
	}

	queue := []*Vertex[T]{node}

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbour := range current.GetNeighbors() {
			if _, ok := m[neighbour.Data]; !ok {
				queue = append(queue, neighbour)
				m[neighbour.Data] = NewVertex(neighbour.Data)
			}
			m[current.Data] = m[neighbour.Data]
		}
	}

	return m[node.Data]
}
