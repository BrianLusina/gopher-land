package courseschedule

const WHITE = 1
const GRAY = 2
const BLACK = 3

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

	for i := 0; i < numCourses; i++ {
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
				if color[neighbour] == WHITE {
					dfs(neighbour)
				} else if color[neighbour] == GRAY {
					isPossible = false
				}
			}
		}

		color[node] = BLACK
	}

	for index := 0; index < numCourses; index++ {
		if color[index] == WHITE {
			dfs(index)
		}
	}

	return isPossible
}
