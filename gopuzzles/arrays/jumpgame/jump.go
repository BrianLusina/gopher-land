package jumpgame

import "gopherland/math/utils"

func jump(nums []int) int {
	size := len(nums)
	// destination is last index
	destination := size - 1

	// record of current coverage, record of last jump index
	currentCoverage := 0
	lastJumpIndex := 0

	// min number of jumps
	minJumps := 0

	// if start index == destination index == 0
	if size == 1 {
		return 0
	}

	// Greedy strategy: extend coverage as long as possible with lazy jump
	for idx := range nums {
		// extend coverage as far as possible
		currentCoverage = utils.Max(currentCoverage, idx+nums[idx])

		// forced to jump (by lazy jump) to extend coverage
		if idx == lastJumpIndex {
			// update last jump index
			lastJumpIndex = currentCoverage

			// update counter of jump by +1
			minJumps++

			// check if destination has been reached
			if currentCoverage >= destination {
				return minJumps
			}
		}
	}

	return minJumps
}
