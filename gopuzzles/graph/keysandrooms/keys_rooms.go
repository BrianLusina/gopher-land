package keysandrooms

// canVisitAllRooms checks if it is possible to visit all rooms from the starting room numbered 0
// Checks if it's possible to visit all rooms if we have only the 1st room as open. The input rooms is a 2D array/
// matrix with all the rooms with each room's number as the index and the value at each index being the list of keys
// for other room numbers.
// Space Complexity: O(n+k) as we are keeping track of visited nodes in a set and k is the number of keys.
// Time Complexity: O(n) as we are traversing all nodes in the graph
// rooms is an adjacency list where room[i] represents the adjacent vertices of the vertex i.
// returns whether it is possible to visit all the vertices/rooms in the adjacency list
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
		// dequeue the vertex from the queue
		vertex := queue[0]
		queue = queue[1:]

		// get all the vertex adjacent neighbors, i.e. rooms we can visit from current room
		neighbors := rooms[vertex]

		for _, neighbor := range neighbors {
			// check if we haven't visited this room
			if _, ok := visitedNodes[neighbor]; !ok {
				// add that we have visited this node
				visitedNodes[neighbor] = true
				// add it for processing
				queue = append(queue, neighbor)
			}
		}
	}

	// At this point, the length of visited nodes should equal to the initial length of rooms. if the length is not equal
	// it means it is not possible to visit all the rooms from the starting room
	return len(visitedNodes) == len(rooms)
}
