package numberofclosedislands

import "testing"

type testCase struct {
	name     string
	grid     [][]int
	expected int
}

var testCases = []testCase{
	{
		name:     "Grid of [[]]",
		grid:     [][]int{},
		expected: 0,
	},
	{
		name:     "Grid of [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]",
		grid:     [][]int{{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}},
		expected: 2,
	},
	{
		name:     "Grid of [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]",
		grid:     [][]int{{0, 0, 1, 0, 0}, {0, 1, 0, 1, 0}, {0, 1, 1, 1, 0}},
		expected: 1,
	},
	{
		name: "Grid of [[1,1,1,1,1,1,1], [1,0,0,0,0,0,1], [1,0,1,1,1,0,1], [1,0,1,0,1,0,1], [1,0,1,1,1,0,1], [1,0,0,0,0,0,1], [1,1,1,1,1,1,1]]",
		grid: [][]int{{1, 1, 1, 1, 1, 1, 1},
			{1, 0, 0, 0, 0, 0, 1},
			{1, 0, 1, 1, 1, 0, 1},
			{1, 0, 1, 0, 1, 0, 1},
			{1, 0, 1, 1, 1, 0, 1},
			{1, 0, 0, 0, 0, 0, 1},
			{1, 1, 1, 1, 1, 1, 1}},
		expected: 2,
	},
}

func TestNumberOfClosedIslands(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := numberOfClosedIslands(tc.grid)
			if actual != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func BenchmarkNumberOfClosedIslands(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark for short mode")
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			numberOfClosedIslands(tc.grid)
		}
	}
}
