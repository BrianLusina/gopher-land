// Package utils contains utility functions
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

// Min returns the minimum between 2 integer values
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
