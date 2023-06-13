// Package utils contains utility functions
package utils

import "golang.org/x/exp/constraints"

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

// Sum returns the sum of a slice of integers
func Sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

// Gcd obtains the Greatest Common Divisor of 2 numbers
func Gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return Gcd(y, x%y)
}

// AbsDiff gets the absolute difference between 2 integers
func AbsDiff[T constraints.Integer](x, y T) T {
	if x < y {
		return y - x
	}
	return x - y
}
