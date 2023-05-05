package utils

// Modulo returns the modulo of n and m.
func Modulo(n, m int) int {
	return ((n % m) + m) % m
}

// Max returns the maximum of 2 integers
func Max(n, m int) int {
	if n > m {
		return n
	}
	return m
}
