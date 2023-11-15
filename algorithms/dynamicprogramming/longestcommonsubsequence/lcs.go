package longestcommonsubsequence

import "gopherland/pkg/utils"

// longestCommonSubsequence is a classic problem. Let dpj be the LCS for string text1 ends at index i and string text2 ends at index j.
// If text1i==text2j, then dpj would be 1+dpj−1. Otherwise, we target the largest LCS if we skip one character from either text1 or text2, i.e.
// dpj=max(dpj,dpj−1).
// Complexity: Where m is the length of text1 and n is the length of text2
// Time complexity: O(m*n)
// Space complexity: O(m*n)
func longestCommonSubsequence(text1 string, text2 string) int {
	text1Len := len(text1)
	text2Len := len(text2)

	dp := make([][]int, text1Len+1)

	for i := 0; i < text1Len+1; i++ {
		dp[i] = make([]int, text2Len+1)
	}

	for i := 1; i < text1Len+1; i++ {
		for j := 1; j < text2Len+1; j++ {
			if text1[i-1] == text2[j-1] {
				// if match found, then store value of previous diagonal element(dp[i - 1][j - 1]) and increase the
				// value by 1 i.e. a new character match is found
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// otherwise, choose maximum of either previous element, either in row(dp[i][j -1]) or column(dp[i][j - 1])
				dp[i][j] = utils.Max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	// dp[text1Length][text2Length] would hold the value of the LCS obtained
	return dp[text1Len][text2Len]
}
