package nonoverlapping

import (
	"testing"
)

type testCase struct {
	name      string
	intervals [][]int
	expected  int
}

var testCases = []testCase{
	{
		name:      "intervals = [[1,2],[2,3],[3,4],[1,3]] should return 1",
		intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
		expected:  1,
	},
	{
		name:      "intervals = [[1,2],[1,2],[1,2]] should return 2",
		intervals: [][]int{{1, 2}, {1, 2}, {1, 2}},
		expected:  2,
	},
	{
		name:      "intervals = [[1,2],[2,3]] should return 0",
		intervals: [][]int{{1, 2}, {2, 3}},
		expected:  0,
	},
}

func TestFindMinRemovalNonOverlapIntervals(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := FindMinRemovalOfNonOverlapIntervals(tc.intervals)
			if actual != tc.expected {
				t.Fatalf("FindMinRemovalOfNonOverlapIntervals(%v) = %d is not equal to expected %d", tc.intervals, actual, tc.expected)
			}
		})
	}
}

func BenchmarkFindMinRemovalOfNonOverlapIntervals(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			FindMinRemovalOfNonOverlapIntervals(tc.intervals)
		}
	}
}
