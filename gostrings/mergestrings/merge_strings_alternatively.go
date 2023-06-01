// Package mergestrings contains functions that are used to merge 2 strings alternatively
package mergestrings

import "gopherland/math/utils"

// MergeAlternatively merges 2 strings into 1 alternatively using 2 pointers
func MergeAlternately(word1 string, word2 string) string {
	mergedWord := ""
	word1Len := len(word1)
	word2Len := len(word2)
	firstPointer := 0
	secondPointer := 0

	for firstPointer < word1Len || secondPointer < word2Len {
		if firstPointer < word1Len {
			mergedWord += string(word1[firstPointer])
			firstPointer++
		}

		if secondPointer < word2Len {
			mergedWord += string(word2[secondPointer])
			secondPointer++
		}
	}

	return mergedWord
}

// MergeAlternatively2 merges 2 strings into 1 alternatively using 1 pointer
func MergeAlternately2(word1 string, word2 string) string {
	mergedWord := ""
	word1Len := len(word1)
	word2Len := len(word2)

	n := utils.Max(word1Len, word2Len)

	for pointer := 0; pointer <= n; pointer++ {
		if pointer < word1Len {
			mergedWord += string(word1[pointer])
		}
		if pointer < word2Len {
			mergedWord += string(word2[pointer])
		}
	}

	return mergedWord
}
