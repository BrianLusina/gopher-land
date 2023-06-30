package maxaveragesubarray

import "gopherland/math/utils"

// findMaxAverage finds the maximum average in sub array of length k
func findMaxAverage(nums []int, k int) float64 {
	if len(nums) <= 1 || len(nums) == k {
		return float64(utils.Sum(nums)) / float64(k)
	}

	currentSum := 0
	for x := 0; x < k; x++ {
		currentSum += nums[x]
	}

	currentMaxSum := currentSum
	for x := k; x < len(nums); x++ {
		currentSum += (nums[x] - nums[x-k])
		currentMaxSum = utils.Max(currentSum, currentMaxSum)
	}

	return float64(currentMaxSum) / float64(k)
}
