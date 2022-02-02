package maxsubarray

import "math"

func maxSubArray(nums []int) int {
	maxSoFar := math.Inf(-1)
	maxEnding := 0

	for _, num := range nums {
		maxEnding += num

		if maxSoFar < float64(maxEnding) {
			maxSoFar = float64(maxEnding)
		}

		if maxEnding < 0 {
			maxEnding = 0
		}
	}

	return int(maxSoFar)
}
