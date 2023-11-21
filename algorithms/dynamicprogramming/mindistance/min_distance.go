package mindistance

import "gopherland/pkg/utils"

func minDistance(word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	}

	if len(word1) == 0 {
		return len(word2)
	}

	m := len(word1)
	n := len(word2)

	dp := make([][]int, m+1)

	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for j := 0; j < n+1; j++ {
		dp[0][j] = j
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var v int
			if word1[i] == word2[j] {
				v = 0
			} else {
				v = 1
			}
			// swap or identical character
			x := dp[i][j] + v

			// # remove last character
			y := dp[i][j+1] + 1

			// # add last character
			z := dp[i+1][j] + 1

			dp[i+1][j+1] = utils.Min(x, utils.Min(y, z))
		}
	}

	return dp[m][n]
}
