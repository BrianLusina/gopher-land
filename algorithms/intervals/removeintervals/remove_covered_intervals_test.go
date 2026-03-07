package removeintervals

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	intervals [][]int
	expected  int
}

var testCases = []testCase{
	{
		intervals: [][]int{
			{1, 5},
			{2, 5},
			{3, 5},
			{4, 5},
		},
		expected: 1,
	},
	{
		intervals: [][]int{
			{1, 3},
			{3, 6},
			{6, 9},
		},
		expected: 3,
	},
	{
		intervals: [][]int{
			{1, 3},
			{4, 6},
			{7, 9},
		},
		expected: 3,
	},
	{
		intervals: [][]int{
			{1, 10},
			{2, 9},
			{3, 8},
			{4, 7},
		},
		expected: 1,
	},
	{
		intervals: [][]int{
			{1, 4},
			{3, 6},
			{2, 8},
		},
		expected: 2,
	},
	{
		intervals: [][]int{
			{1, 2},
			{1, 4},
			{3, 4},
		},
		expected: 1,
	},
	{
		intervals: [][]int{
			{1, 5},
			{2, 3},
			{4, 6},
		},
		expected: 2,
	},
}

func TestRemoveCoveredIntervals(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("RemoveCoveredIntervals(%v)", tc.intervals), func(t *testing.T) {
			actual := RemoveCoveredIntervals(tc.intervals)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkRemoveCoveredIntervals(b *testing.B) {
	for b.Loop() {
		for _, tc := range testCases {
			_ = RemoveCoveredIntervals(tc.intervals)
		}
	}
}
