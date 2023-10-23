package lettercombinations

import "strings"

func letterCombinations(n int) []string {
	result := []string{}
	if n == 0 {
		return result
	}

	var dfs func(int, []string)
	dfs = func(startIndex int, path []string) {
		if startIndex == n {
			result = append(result, strings.Join(path, ""))
			return
		}

		for _, char := range []string{"a", "b"} {
			path = append(path, char)
			dfs(startIndex+1, path)
			path = path[:len(path)-1]
		}
	}

	dfs(0, []string{})

	return result
}
