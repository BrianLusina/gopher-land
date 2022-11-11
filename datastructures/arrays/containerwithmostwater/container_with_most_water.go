package containerwithmostwater

import "math"

// maxArea gets the maximum area of a container that can contain water in linear time
func maxArea(heights []int) int {
	result := 0

	left, right := 0, len(heights)-1

	for left < right {
		area := (right - left) * int(math.Min(float64(heights[left]), float64(heights[right])))
		result = int(math.Max(float64(result), float64(area)))

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return result
}
