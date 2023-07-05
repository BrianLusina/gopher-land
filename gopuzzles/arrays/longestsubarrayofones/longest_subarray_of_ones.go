package longestsubarrayofones

import "gopherland/math/utils"

// longestSubarray finds the longest subarray of ones from a binary array after deleting 1 element
func longestSubarray(nums []int) int {
	longestWindow := 0
	left := 0
	lastZero := -1

	for right := range nums {
		if nums[right] == 0 {
			left = lastZero + 1
			lastZero = right
		}

		longestWindow = utils.Max(longestWindow, right-left)
	}

	return longestWindow
}
