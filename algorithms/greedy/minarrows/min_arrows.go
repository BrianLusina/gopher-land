package minarrows

import "sort"

// findMinArrowShots finds the minimum number of arrows to burst all balloons
// The answer is at least 1. First we sort the balloons by the end coordinate. Set the first end coordinate as
// current. Iterate over all balloons to check if the balloon starts after current. If so, increase answer by 1 and set
// current = right.
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	count := 1
	current := points[0][1]

	for _, point := range points {
		left := point[0]
		right := point[1]

		if current < left {
			count++
			current = right
		}
	}

	return count
}
