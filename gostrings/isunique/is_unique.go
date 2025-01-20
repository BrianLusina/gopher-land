package isunique

// IsUnique checks if a string has all unique characters
func IsUnique(inputStr string) bool {
	// Create a map to store the frequency of each character
	charMap := make(map[rune]bool)

	// Iterate through the string
	for _, char := range inputStr {
		// If the character is already in the map, return false
		if charMap[char] {
			return false
		}

		// Otherwise, add the character to the map
		charMap[char] = true
	}

	// If the loop completes, return true
	return true
}
