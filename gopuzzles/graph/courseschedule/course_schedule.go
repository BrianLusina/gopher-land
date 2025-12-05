package courseschedule

const WHITE = 1
const GRAY = 2
const BLACK = 3

// CanFinish checks if it is possible to finish the number of courses given a slice of prerequisite courses to do
func CanFinish(numCourses int, prerequisites [][]int) bool {
	adjacencyList := map[int][]int{}
	color := make(map[int]int)

	for _, prerequisite := range prerequisites {
		source := prerequisite[1]
		destination := prerequisite[0]

		neighbours, ok := adjacencyList[source]

		if ok {
			neighbours = append(neighbours, destination)
			adjacencyList[source] = neighbours
		} else {
			adjacencyList[source] = []int{destination}
		}
	}

	isPossible := true

	for i := range numCourses {
		color[i] = WHITE
	}

	var dfs func(node int)

	dfs = func(node int) {
		if !isPossible {
			return
		}

		color[node] = GRAY

		neighbours, ok := adjacencyList[node]

		if ok {
			for _, neighbour := range neighbours {
				switch color[neighbour] {
				case WHITE:
					dfs(neighbour)
				case GRAY:
					isPossible = false
				}
			}
		}

		color[node] = BLACK
	}

	for index := range numCourses {
		if color[index] == WHITE {
			dfs(index)
		}
	}

	return isPossible
}
