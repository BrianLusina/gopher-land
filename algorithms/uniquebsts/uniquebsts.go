package uniquebsts

func NumTrees(n int) int {
	dp := make([]int, n+1)

	dp[0], dp[1] = 1, 1

	for x := 2; x < n+1; x++ {
		for y := 1; y < x+1; y++ {
			dp[x] = dp[x] + (dp[x-y] * dp[y-1])
		}
	}

	return dp[n]
}
