package maxsubarray

import (
	"math"
)

// MaxSubarray finds the maximum sub array of a slice of int
func MaxSubarray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	bestSum := math.Inf(-1)
	currentSum := 0

	for _, number := range numbers {
		currentSum = int(math.Max(float64(number), float64(currentSum)+float64(number)))
		bestSum = math.Max(bestSum, float64(currentSum))
	}
	return int(bestSum)
}
