package hindex

import "testing"

type testCase struct {
	citations []int
	expected  int
}

var testCases = []testCase{
	{
		citations: []int{3, 0, 6, 1, 5},
		expected:  3,
	},
	{
		citations: []int{1, 3, 1},
		expected:  1,
	},
	{
		citations: []int{9, 7, 6, 2, 1},
		expected:  3,
	},
}

func TestHIndex(t *testing.T) {
	for _, tc := range testCases {
		actual := hIndex(tc.citations)
		if actual != tc.expected {
			t.Fatalf("hIndex(%v) = %d, expected %d", tc.citations, actual, tc.expected)
		}
	}
}

func BenchmarkHIndex(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := hIndex(tc.citations)
			if actual != tc.expected {
				b.Fatalf("hIndex(%v) = %d, expected %d", tc.citations, actual, tc.expected)
			}
		}
	}
}
