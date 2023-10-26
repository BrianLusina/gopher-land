package dominotrominotiling

const MOD = 1000000007

func numTilings(n int) int {
	if n < 3 {
		return n
	}

	dp := [][3]int{}

	for i := 0; i <= n; i++ {
		dp = append(dp, [3]int{})
	}

	dp[0][0] = 1
	dp[1][0] = 1
	dp[1][1] = 1
	dp[1][2] = 1

	for i := 2; i <= n; i++ {
		dp[i][0] = (dp[i-1][0] +
			dp[i-2][0] +
			dp[i-2][1] +
			dp[i-2][2]) % MOD

		dp[i][1] = (dp[i-1][0] +
			dp[i-1][2]) % MOD

		dp[i][2] = (dp[i-1][0] +
			dp[i-1][1]) % MOD
	}

	return dp[n][0]
}
