package findtownjudge

// FindJudge
// Consider trust as a graph, all pairs are directed edge.
// The point with in-degree - out-degree = N - 1 become the judge.
//
// Explanation:
//
// Count the degree, and check at the end.
//
// Time Complexity:
// Time O(T + N), space O(N)
func FindJudge(n int, trust [][]int) int {
	count := make([]int, n+1)

	for _, t := range trust {
		count[t[0]]--
		count[t[1]]++
	}

	for i := 1; i < len(count); i++ {
		if count[i] == n-1 {
			return i
		}
	}
	return -1
}
