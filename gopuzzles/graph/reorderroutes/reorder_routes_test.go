package reorderroutes

import "testing"

type testCase struct {
	n           int
	connections [][]int
	expected    int
}

var testCases = []testCase{
	{
		connections: [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}},
		n:           6,
		expected:    3,
	},
	{
		connections: [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}},
		n:           5,
		expected:    2,
	},
	{
		connections: [][]int{{1, 0}, {2, 0}},
		n:           3,
		expected:    0,
	},
}

func TestMinReorder(t *testing.T) {
	for _, tc := range testCases {
		actual := minReorder(tc.n, tc.connections)
		if actual != tc.expected {
			t.Fatalf("minReorder(%d, %v) = %d, expected=%d", tc.n, tc.connections, actual, tc.expected)
		}
	}
}

func BenchmarkMinReorder(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := minReorder(tc.n, tc.connections)
			if actual != tc.expected {
				b.Fatalf("minReorder(%d, %v) = %d, expected=%d", tc.n, tc.connections, actual, tc.expected)
			}
		}
	}
}
