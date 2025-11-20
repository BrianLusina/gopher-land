// Package utils contains utility functions
package utils

import (
	"gopherland/pkg/types"

	"golang.org/x/exp/constraints"
)

// Modulo returns the modulo of n and m.
func Modulo(n, m int) int {
	return ((n % m) + m) % m
}

func DivMod(n, m int) (int, int) {
	rem := Modulo(n, m)
	n /= m
	return n, rem
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

// Sum returns the sum of a slice of numbers
func Sum[T types.Number](numbers []T) T {
	var result T
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

// Abs gets the absolute value of an integer
func Abs(x int) int {
	if x < 1 {
		return x * -1
	}
	return x
}

// sumOfSquaredDigits returns the sum of the squares of the digits of n.
// For example, sumOfSquaredDigits(19) returns 82 (1^2 + 9^2).
func SumOfSquaredDigits(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}
