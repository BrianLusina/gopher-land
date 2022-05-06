package graphs

// CloneGraph clones a graph and returns a new graph
// this assumes the graph is connected and undirected
func CloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	m := map[any]*Node{
		node.Data: NewNode(node.Data),
	}

	queue := []*Node{node}

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbour := range current.Neighbors {
			if _, ok := m[neighbour.Data]; !ok {
				queue = append(queue, neighbour)
				m[neighbour.Data] = NewNode(neighbour.Data)
			}
			m[current.Data] = m[neighbour.Data]
		}
	}

	return m[node.Data]
}
