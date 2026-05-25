package wordpattern

import (
	"fmt"
	"strings"
)

// wordPattern checks if a pattern matches a string following a bijective mapping. Each character in pattern should map to exactly one word in s, and vice versa.
func wordPattern(pattern string, s string) bool {
	// Split the input string s into words using space as a delimiter.
	words := strings.Split(s, " ")

	// If the number of characters in the pattern does not match the number of words, return false immediately.
	if len(pattern) != len(words) {
		return false
	}

	patternToWord := make(map[byte]string) // Map to store the mapping from pattern characters to words.
	wordToPattern := make(map[string]byte) // Map to store the mapping from words to pattern characters.

	// Iterate through the pattern and corresponding words simultaneously.
	for index := 0; index < len(pattern); index++ {
		patternChar := pattern[index] // Current character in the pattern.
		word := words[index]          // Corresponding word in the string.

		// Check if the current pattern character has an existing mapping.
		if mappedWord, exists := patternToWord[patternChar]; exists {
			// If it exists but does not match the current word, return false.
			if mappedWord != word {
				return false
			}
		} else {
			// If there is no existing mapping for the pattern character, check if the current word is already mapped to a different pattern character.
			if mappedChar, exists := wordToPattern[word]; exists {
				if mappedChar != patternChar {
					return false
				}
			} else {
				// If there is no existing mapping for the word, create new mappings in both directions.
				patternToWord[patternChar] = word
				wordToPattern[word] = patternChar
			}
		}
	}

	// If all characters and words match the bijective mapping, return true.
	return true
}

// wordPattern2 checks if a pattern matches a string following a bijective mapping. Each character in pattern should map to exactly one word in s, and vice versa.
func wordPattern2(pattern string, s string) bool {
	mapIndex := make(map[string]int) // Map to store the index mapping for both characters and words

	// Split the input string s into words using space as a delimiter.
	words := strings.Split(s, " ")

	// If the number of characters in the pattern does not match the number of words, return false immediately.
	if len(pattern) != len(words) {
		return false
	}

	// Iterate through the pattern and corresponding words simultaneously.
	for index := range words {
		patternChar := pattern[index] // Current character in the pattern.
		word := words[index]          // Corresponding word in the string.

		charKey := fmt.Sprintf("char_%c", patternChar) // Unique key for the pattern character.
		wordKey := fmt.Sprintf("word_%s", word)        // Unique key for the word.

		// If this character hasn't been seen before, store its index
		if _, exists := mapIndex[charKey]; !exists {
			mapIndex[charKey] = index
		}

		// If this word hasn't been seen before, store its index
		if _, exists := mapIndex[wordKey]; !exists {
			mapIndex[wordKey] = index
		}

		//  If the indices for the character and word don't match,it means they were mapped inconsistently — return False
		if mapIndex[charKey] != mapIndex[wordKey] {
			return false
		}
	}

	// If all characters and words match the bijective mapping, return true.
	return true
}
