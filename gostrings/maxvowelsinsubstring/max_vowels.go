package maxvowelsinsubstring

import "gopherland/math/utils"

func maxVowels(s string, k int) int {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	currentMaxVowels := 0
	maxLength := 0

	for x := 0; x < len(s); x++ {
		if vowels[s[x]] {
			currentMaxVowels++
		}
		if x >= k {
			if vowels[s[x-k]] {
				currentMaxVowels--
			}
		}

		maxLength = utils.Max(currentMaxVowels, maxLength)
	}

	return maxLength
}
