package constructrectangle

import "math"

func constructRectangle(area int) []int {
	w := int(math.Sqrt(float64(area)))

	for area%w != 0 {
		w -= 1
	}

	return []int{int(area / w), w}
}
