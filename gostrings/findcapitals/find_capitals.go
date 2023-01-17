package findcapitals

import (
	"unicode"
)

func findCapitals(word string) []int {
	caps := []int{}
	for idx, char := range word {
		if unicode.IsUpper(char) {
			caps = append(caps, idx)
		}
	}

	return caps
}
