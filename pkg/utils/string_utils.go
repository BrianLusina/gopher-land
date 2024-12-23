package utils

// isAlphanumeric returns true if the character is an alphanumeric character
// this uses the ASCII table to quickly find the position of the character. Since the characters A-Z are contiguous, simply checking whether this
// character falls within the range will determine whether this character is alphanumeric or not. Same case applies for a-z & 0-9
func IsAlphanumeric(char byte) bool {
	return ('A' <= char && char <= 'Z') || ('a' <= char && char <= 'z') || ('0' <= char && char <= '9')
}

// Reverse reverses a given string
func Reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
