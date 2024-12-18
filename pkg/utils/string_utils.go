package utils

// isAlphanumeric returns true if the character is an alphanumeric character
// this uses the ASCII table to quickly find the position of the character. Since the characters A-Z are contiguous, simply checking whether this
// character falls within the range will determine whether this character is alphanumeric or not. Same case applies for a-z & 0-9
func IsAlphanumeric(char byte) bool {
	return ('A' <= char && char <= 'Z') || ('a' <= char && char <= 'z') || ('0' <= char && char <= '9')
}
