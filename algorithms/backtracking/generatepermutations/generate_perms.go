package generatepermutations

// generatePermutations Generates permutations of the given letters and returns a list of all permutations.
// Time Complexity We have n letters to choose in the first level, n - 1 choices in the second level and so on therefore the number of strings
// we generate is n * (n - 1) * (n - 2) * ... * 1, or O(n!).
// Since each string has length n, generating all the strings requires O(n * n!) time.
//
// Space Complexity The total space complexity is given by the amount of space required by the strings we're constructing.
// Like the time complexity, the space complexity is also O(n * n!).
func generatePermutations(letters string) []string {
	var dfs func(startIndex int, path []rune, used []bool, result *[]string)
	dfs = func(startIndex int, path []rune, used []bool, result *[]string) {
		if startIndex == len(letters) {
			p := string(path)
			*result = append(*result, p)
			return
		}

		for idx, letter := range letters {
			// skip used letters
			if used[idx] {
				continue
			}

			// add letter to permutation, mark letter as used
			path = append(path, letter)
			used[idx] = true
			dfs(startIndex+1, path, used, result)

			// remove letter from permutation, mark letter as unused
			path = path[:len(path)-1]
			used[idx] = false
		}
	}

	res := []string{}
	used := make([]bool, len(letters))
	dfs(0, []rune{}, used, &res)
	return res
}
