package containerwithmostwater

import "testing"

type testCase struct {
	heights  []int
	expected int
}

var testCases = []testCase{
	{
		heights:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		expected: 49,
	},
	{
		heights:  []int{1, 1},
		expected: 1,
	},
}

func TestContainerWithMostWater(t *testing.T) {
	for _, tc := range testCases {
		actual := maxArea(tc.heights)
		if actual != tc.expected {
			t.Fatalf("maxArea(%v) = %d, expected= %d", tc.heights, actual, tc.expected)
		}
	}
}

func BenchmarkContainerWithMostWater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			maxArea(tc.heights)
		}
	}
}
