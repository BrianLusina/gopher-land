package mergeintervals

import (
	"math"
	"sort"
)

// type ByStartInterval [][]int

// func (a ByStartInterval) Len() int           { return len(a) }
// func (a ByStartInterval) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a ByStartInterval) Less(i, j int) bool { return a[i][0] < a[j][0] }

func MergeIntervals(intervals [][]int) (merged [][]int) {
	// sort.Sort(ByStartInterval(intervals))

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for _, interval := range intervals {
		if len(merged) == 0 || merged[len(merged)-1][1] < interval[0] {
			merged = append(merged, interval)
		} else {
			merged[len(merged)-1][1] = int(math.Max(float64(merged[len(merged)-1][1]), float64(interval[1])))
		}
	}

	return
}
