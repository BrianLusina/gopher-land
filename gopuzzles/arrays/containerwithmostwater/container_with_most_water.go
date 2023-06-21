package containerwithmostwater

import (
	"gopherland/math/utils"
)

// maxArea gets the maximum area of a container that can contain water in linear time
func maxArea(heights []int) int {
	maxPossibleArea := 0

	left, right := 0, len(heights)-1

	for left < right {
		width := right - left
		area := width * utils.Min(heights[left], heights[right])
		maxPossibleArea = utils.Max(maxPossibleArea, area)

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return maxPossibleArea
}

// maxAreaBruteForce finds the maximum area that can be obtained from a slice of integers
func maxAreaBruteForce(heights []int) int {
	maxarea := 0

	for left := 0; left < len(heights); left++ {
		for right := left + 1; right < len(heights); right++ {
			width := right - left
			maxarea = utils.Max(maxarea, utils.Min(heights[left], heights[right])*width)
		}
	}

	return maxarea
}
