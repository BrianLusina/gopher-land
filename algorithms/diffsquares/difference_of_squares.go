package diffsquares

import (
	"math"
)

// SquareOfSum returns the square of the sum of N natural numbers
func SquareOfSum(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return int(math.Pow(float64(total), 2))
}

// SumOfSquares returns the sum of the squares of the first N natural numbers
func SumOfSquares(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += int(math.Pow(float64(i), 2))
	}
	return total
}

// Difference returns the difference between the square of the sum of the first N natural numbers & sum of the squares of the first N natural numbers
func Difference(n int) int {
	sqOfSum := SquareOfSum(n)
	sumOfSq := SumOfSquares(n)

	return sqOfSum - sumOfSq
}
