package keysandrooms

// canVisitAllRooms checks if it is possible to visit all rooms from the starting room numbered 0
func canVisitAllRooms(rooms [][]int) bool {
	// the starting point in the graph
	startingNode := 0

	// keep track of nodes not yet visited/processed
	queue := []int{}
	queue = append(queue, startingNode)

	// keeps track of nodes/rooms already visited
	visitedNodes := map[int]bool{}
	visitedNodes[startingNode] = true

	for len(queue) != 0 {
		vertex := queue[0]
		queue = queue[1:]

		neighbors := rooms[vertex]

		for _, neighbor := range neighbors {
			if _, ok := visitedNodes[neighbor]; !ok {
				visitedNodes[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return len(visitedNodes) == len(rooms)
}
