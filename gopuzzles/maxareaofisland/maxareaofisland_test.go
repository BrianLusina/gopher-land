package maxareaofisland

import "testing"

type testCase struct {
	grid     [][]int
	expected int
}

var testCases = []testCase{
	{
		grid: [][]int{
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		},
		expected: 6,
	},
	{
		grid:     [][]int{{0, 0, 0, 0, 0, 0, 0, 0}},
		expected: 0,
	},
}

func TestMaxAreaOfIsland(t *testing.T) {
	for _, test := range testCases {
		actual := maxAreaOfIsland(test.grid)
		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
		t.Logf("Passed %d", test.grid)
	}
}

func BenchmarkMaxAreaOfIsland(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			maxAreaOfIsland(test.grid)
		}
	}
}
