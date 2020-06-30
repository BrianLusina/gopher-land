package etl

import (
	"strings"
)

// Transform takes in legacy scores with map of int to array of strings and transforms them
// to a map of strings to int
func Transform(legacyScores map[int][]string) (newScores map[string]int) {
	newScores = make(map[string]int)
	for index, letters := range legacyScores {
		for _, letter := range letters {
			letterLower := strings.ToLower(letter)
			newScores[letterLower] = index
		}
	}
	return
}
