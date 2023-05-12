package nonoverlapping

import (
	"math"
	"sort"
)

// FindMinRemovalOfNonOverlapIntervals finds the minimum number of non overlap intervals that
// should be removed
func FindMinRemovalOfNonOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	currentMinIntervalEnd, count := math.Inf(-1), 0

	for _, interval := range intervals {
		currentStart, currentEnd := interval[0], interval[1]

		if currentStart >= int(currentMinIntervalEnd) {
			currentMinIntervalEnd = float64(currentEnd)
		} else {
			count++
		}
	}

	return count
}
