package rottingoranges

import (
	"fmt"
	"testing"
)

type testCase struct {
	grid     [][]int
	expected int
}

var testCases = []testCase{
	{
		grid:     [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
		expected: 4,
	},
	{
		grid:     [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
		expected: -1,
	},
	{
		grid:     [][]int{{0, 2}},
		expected: 0,
	},
	{
		grid:     [][]int{{2, 1, 1}, {1, 1, 1}, {0, 1, 2}},
		expected: 2,
	},
}

func TestRottingOranges(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("orangesRotting(%v)", tc.grid), func(t *testing.T) {
			actual := orangesRotting(tc.grid)
			if actual != tc.expected {
				t.Fatalf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func BenchmarkRottingOranges(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("orangesRotting(%v)", tc.grid), func(b *testing.B) {
				actual := orangesRotting(tc.grid)
				if actual != tc.expected {
					b.Fatalf("expected %d, got %d", tc.expected, actual)
				}
			})
		}
	}
}
