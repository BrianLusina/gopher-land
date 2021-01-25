package pangram

import "strings"

// IsPangram checks if a sentence is a pangram & returns true
func IsPangram(sentence string) bool {
	runeCounts := countRunes(strings.ToLower(sentence))

	for asc := 'a'; asc <= 'z'; asc++ {
		if runeCounts[asc] == 0 {
			return false
		}
	}
	return true
}

// countRunes builds a map of runes & how often they appear in a string
func countRunes(s string) map[rune]int {
	m := make(map[rune]int)

	for _, runeVal := range s {
		m[runeVal]++
	}
	return m
}
