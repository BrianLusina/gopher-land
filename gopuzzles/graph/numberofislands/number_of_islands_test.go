package numberofislands

import "testing"

type testCases struct {
	grid     [][]byte
	expected int
}

var tests = []testCases{
	{
		grid:     [][]byte{},
		expected: 0,
	},
	{
		grid: [][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		},
		expected: 1,
	},
	{
		grid: [][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		},
		expected: 3,
	},
}

func TestNumberOfIslands(t *testing.T) {
	for _, test := range tests {
		actual := NumberOfIslands(test.grid)
		if actual != test.expected {
			t.Errorf("FAIL: \nNumberOfIslands(%v)\nreturned %d, expected %d.", test.grid, actual, test.expected)
		}
		t.Logf("PASS: %s", test.grid)
	}
}

func BenchmarkNumberOfIslands(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			NumberOfIslands(test.grid)
		}
	}
}
