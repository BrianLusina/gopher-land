package minnumberofvertices

// FindSmallesSetOfVertices finds the smallest number of vertices that can be reached from source nodes in a
// Directed Acyclic Graph
func FindSmallestSetOfVertices(n int, edges [][]int) (vertices []int) {
	inDegrees := make([]int, n)

	for _, edge := range edges {
		inDegrees[edge[1]]++
	}

	for x := 0; x < n; x++ {
		if inDegrees[x] == 0 {
			vertices = append(vertices, x)
		}
	}

	return
}
