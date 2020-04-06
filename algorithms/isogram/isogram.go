package isogram

import "strings"

// IsIsogram takes in a word and checks if it has repeating letters
func IsIsogram(word string) bool {

	wordLower := strings.ToLower(word)

	for index, letter := range wordLower {
		if letter != '-' && letter != ' ' && index < strings.LastIndex(wordLower, string(letter)) {
			return false
		}
	}

	return true
}
