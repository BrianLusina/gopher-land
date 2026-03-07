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

	// Sort by start time in ascending order and then by end time in descending order
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			// Sort by start time ascending
			return intervals[i][0] < intervals[j][0]
		}

		// If start times are equal, sort by end time descending
		return intervals[i][1] > intervals[j][1]
	})

	maxEndSeen := -math.MaxInt
	count := 0

	for _, interval := range intervals {
		current_end := interval[1]
		if current_end > maxEndSeen {
			count += 1
			maxEndSeen = current_end
		}
	}

	return count
}
