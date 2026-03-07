package removeintervals

import (
	"math"
	"sort"
)

// RemoveCoveredIntervals removes intervals that have already been covered by other intervals and returns the count of non-covered intervals remaining
func RemoveCoveredIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// create a copy of the intervals being passed in to avoid side effects on the caller side. This will incur a space complexity cost of O(n)
	sorted := make([][]int, len(intervals))
	copy(sorted, intervals)

	// Sort by start time in ascending order and then by end time in descending order
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i][0] != sorted[j][0] {
			// Sort by start time ascending
			return sorted[i][0] < sorted[j][0]
		}

		// If start times are equal, sort by end time descending
		return sorted[i][1] > sorted[j][1]
	})

	maxEndSeen := -math.MaxInt
	count := 0

	for _, interval := range sorted {
		current_end := interval[1]
		if current_end > maxEndSeen {
			count += 1
			maxEndSeen = current_end
		}
	}

	return count
}
