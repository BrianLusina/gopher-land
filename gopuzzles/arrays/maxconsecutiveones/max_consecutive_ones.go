package maxconsecutiveones

func longestOnes(nums []int, k int) int {
	left, right := 0, 0

	for right < len(nums) {
		// if we encounter a 0 then we decrement k
		if nums[right] == 0 {
			k--
		}

		// if k < 0 then we need to move the left part of the window forward
		// to try and remove the extra 0's
		if k < 0 {
			// if the left one was zero then we adjust K
			if nums[left] == 0 {
				k++
			}
			// regardless of whether we had a 1 or a 0 we can move left side by 1
			// if we keep seeing 1's the window still keeps moving as-is

			left++
		}

		right++
	}

	return right - left
}

// findMaxConsecutiveOnes returns the maximum number of consecutive 1s in the binary array nums.
// The most straightforward way to solve this is to use a single pass through the array, keeping track of two values
// as we go. First, we need a counter for the current streak of consecutive 1s we're seeing. Second, we need to
// remember the maximum streak we've encountered so far.
//
// As we examine each element, if we see a 1, we increment our current streak counter. If we see a 0, that breaks our
// streak, so we reset the counter to 0. Importantly, every time we update our current streak, we check whether it's
// larger than our maximum and update the maximum if needed.
//
// Time complexity is O(n) where n is the length of the input array
// Space complexity is O(1) as no extra space is required
// Reference: https://leetcode.com/problems/max-consecutive-ones/
func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxConsecutiveOnes := 0
	currentConsecutiveOnes := 0

	for _, num := range nums {
		if num == 1 {
			currentConsecutiveOnes += 1
			if maxConsecutiveOnes < currentConsecutiveOnes {
				maxConsecutiveOnes = currentConsecutiveOnes
			}
		} else {
			currentConsecutiveOnes = 0
		}
	}
	return maxConsecutiveOnes
}
