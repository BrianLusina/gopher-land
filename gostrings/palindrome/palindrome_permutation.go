package palindrome

import "strings"

// isPalindromePermutation checks if a string is a palindrome permutation.
func isPalindromePermutation(input_str string) bool {
	// normalize the input string by converting it to lowercase
	normalizedInputStr := strings.ToLower(strings.ReplaceAll(input_str, " ", ""))

	// create a map to store the frequency of each character
	frequency := make(map[rune]int)

	// iterate through the input string
	for _, char := range normalizedInputStr {
		// increment the frequency of the character
		frequency[char]++
	}

	// create a variable to store the number of odd frequencies
	oddCount := 0

	// iterate through the frequency map
	for _, count := range frequency {
		// if the count is odd
		if count%2 != 0 && oddCount == 0 {
			// increment the odd count
			oddCount++
		} else if count%2 != 0 && oddCount != 0 {
			// if there is more than one character with an odd frequency
			// then it is not a palindrome permutation
			return false
		}
	}

	// if there is more than one character with an odd frequency
	// then it is not a palindrome permutation
	return true
}
