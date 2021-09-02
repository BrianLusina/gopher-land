package findifpathexists

func validPath(n int, edges [][]int, start int, end int) bool {
	adjacency_list := make(map[int][]int)
	for _, e := range edges {
		adjacency_list[e[0]] = append(adjacency_list[e[0]], e[1])
		adjacency_list[e[1]] = append(adjacency_list[e[1]], e[0])
	}

	visited := make(map[int]bool)
	stack := []int{start}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node == end {
			return true
		}

		if !visited[node] {
			visited[node] = true
			stack = append(stack, adjacency_list[node]...)
		}
	}
	return false
}
