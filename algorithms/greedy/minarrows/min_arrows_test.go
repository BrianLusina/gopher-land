package minarrows

import "testing"

type testCase struct {
	points   [][]int
	expected int
}

var testCases = []testCase{
	{
		points:   [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
		expected: 2,
	},
	{
		points:   [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
		expected: 4,
	},
	{
		points:   [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
		expected: 2,
	},
}

func TestMinArrowShots(t *testing.T) {
	for _, tc := range testCases {
		actual := findMinArrowShots(tc.points)
		if actual != tc.expected {
			t.Fatalf("findMinArrowShots(points=%v) = %d, expected=%d", tc.points, actual, tc.expected)
		}
	}
}

func BenchmarkMinArrowShots(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := findMinArrowShots(tc.points)
			if actual != tc.expected {
				b.Fatalf("findMinArrowShots(points=%v) = %d, expected=%d", tc.points, actual, tc.expected)
			}
		}
	}
}
