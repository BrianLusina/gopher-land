package numberofprovinces

type unionFind struct {
	root []int
	rank []int
	size int
}

func newUnionFind(size int) *unionFind {
	rank := make([]int, size)
	root := make([]int, size)

	for i := 0; i < size; i++ {
		root[i] = i
		rank[i] = 1
	}

	return &unionFind{
		root: root,
		size: size,
		rank: rank,
	}
}

func (uf *unionFind) find(x int) int {
	if x == uf.root[x] {
		return x
	}
	uf.root[x] = uf.find(uf.root[x])
	return uf.root[x]
}

func (uf *unionFind) union(x, y int) {
	xRoot := uf.find(x)
	yRoot := uf.find(y)

	if xRoot == yRoot {
		return
	}

	if uf.rank[xRoot] < uf.rank[yRoot] {
		uf.root[xRoot] = yRoot
	} else if uf.rank[xRoot] > uf.rank[yRoot] {
		uf.root[yRoot] = xRoot
	} else {
		uf.root[yRoot] = xRoot
		uf.rank[xRoot]++
	}
	uf.size--
}

func numberOfProvinces(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	nrows := len(grid)
	uf := newUnionFind(nrows)

	for row := 0; row < nrows; row++ {
		for col := row; col < nrows; col++ {
			if grid[row][col] == 1 {
				uf.union(row, col)
			}
		}
	}

	return uf.size
}

// numberOfProvincesDfs finds the number of connected components(graphs) in a given adjacency list is_connected.
// Uses DFS traversal algorithm to find the connected components or the number of graphs from the provided 2D array.
// Complexity:
// Time Complexity: O(n^2). Initializing the visited array takes O(n) time. The dfs function visits each node once,
// which takes O(n) time because there are n nodes in total. From each node, we iterate over all possible edges using graph node which takes O(n) time for each visited node.
// As a result, it takes a total of O(n^2) time to visit all the nodes and iterate over its edges.

// Space Complexity: O(n) visited array takes O(n) space Recursion call stack used by dfs can have no more than n elements in the worst case. It would take up to O(n) space in that case.

// isConnected - Array: 2D Array denoting an adjacency list where the index is a city and the values at that index denote interconnectivity with 1 showing a direct connection and 0 showing no connection
// Int: Number of connected components or number of graphs from the list
func numberOfProvincesDfs(isConnected [][]int) int {
	// number of cities is the length of the rows from the grid
	cities := len(isConnected)

	// keep track of number of provinces
	provinces := 0

	// keep track of visited cities
	visited := make([]bool, cities)

	var dfs func(node int, graph [][]int, visit []bool)

	dfs = func(node int, graph [][]int, visit []bool) {
		// mark node as visited
		visit[node] = true

		for i := 0; i < len(graph); i++ {
			// get the city's neighbour to check if it can be reached, in this case if the value is 1, we can reach
			// city's neighbour, i.e, they are connected
			neighbour := graph[node][i]
			// Checks if the neighbour has already been visited
			neighbourVisited := visit[i]

			if neighbour == 1 && !neighbourVisited {
				dfs(i, graph, visit)
			}
		}
	}

	for city := 0; city < cities; city++ {
		// if city has not been visited
		if !visited[city] {
			// increment number of provinces
			provinces++
			// perform a dfs from the city to other cities
			dfs(city, isConnected, visited)
		}
	}

	return provinces
}
