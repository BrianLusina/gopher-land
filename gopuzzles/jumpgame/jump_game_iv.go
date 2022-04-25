package jumpgame

func buildGraph(arr []int) map[int][]int {
	graph := make(map[int][]int, len(arr))

	for idx, value := range arr {
		if _, ok := graph[value]; ok {
			graph[value] = append(graph[value], idx)
		} else {
			graph[value] = []int{idx}
		}
	}
	return graph
}

func minJumps(arr []int) int {
	size := len(arr)
	if size <= 1 {
		return 0
	}

	graph := buildGraph(arr)
	current := []int{0}
	visited := make(map[int]bool)
	step := 0

	for len(current) > 0 {
		next := make([]int, 0)
		for _, node := range current {
			if node == size-1 {
				return step
			}

			for _, child := range graph[arr[node]] {
				if _, ok := visited[child]; !ok {
					next = append(next, child)
					visited[child] = true
				}
			}

			graph[arr[node]] = nil

			for _, child := range []int{node - 1, node + 1} {
				if child >= 0 && child < size && !visited[child] {
					next = append(next, child)
					visited[child] = true
				}
			}
		}
		current = next
		step++
	}
	return -1
}
