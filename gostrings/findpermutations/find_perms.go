package findpermutations

import "strings"

func permutate(word string, start, end int) string {
	chars := strings.Split(word, "")

	if start == end {
		return strings.Join(chars, "")
	}

	for i := start; i <= end; i++ {
		chars[start], chars[i] = chars[i], chars[start]
		permutate(chars[start+1], start+1, end)
		chars[start], chars[i] = chars[i], chars[start]
	}

	return ""
}

func FindPermutations(word string) []string {
	var result []string

	// l := len(word)
	// return permutate(word, 0, l)

	if len(word) == 0 {
		return []string{}
	}

	if len(word) == 1 {
		return []string{word}
	}

	for i := 0; i < len(word); i++ {
		letter := string(word[i])
		rest := word[:i] + word[i+1:]
		for _, perm := range FindPermutations(rest) {
			result = append(result, letter+perm)
		}
	}

	return result
}
