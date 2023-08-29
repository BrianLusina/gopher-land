package trappedrain

import "gopherland/math/utils"

func trap(height []int) int {
	if len(height) == 0 || len(height) == 1 {
		return 0
	}

	left, right := 0, len(height)-1
	max, min := 0, 0
	rainWater := 0

	for left <= right {
		min = utils.Min(height[left], height[right])
		max = utils.Max(max, min)

		rainWater += max - min

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return rainWater
}
