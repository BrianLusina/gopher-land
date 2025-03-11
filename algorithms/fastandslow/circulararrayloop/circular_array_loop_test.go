package circulararrayloop

import (
	"fmt"
	"testing"
)

type testCase struct {
	nums     []int
	expected bool
}

var testCases = []testCase{
	{
		nums:     []int{3, 1, 2},
		expected: true,
	},
	{
		nums:     []int{-2, -1, -3},
		expected: true,
	},
	{
		nums:     []int{2, 1, -1, -2},
		expected: false,
	},
	{
		nums:     []int{3, -3, 1, 1},
		expected: true,
	},
	{
		nums:     []int{1, 3, -2, -4, 1},
		expected: true,
	},
	{
		nums:     []int{5, 4, -2, -1, 3},
		expected: false,
	},
	{
		nums:     []int{1, 2, -3, 3, 4, 7, 1},
		expected: true,
	},
	{
		nums:     []int{3, 3, 1, -1, 2},
		expected: true,
	},
}

// TestCircularArrayLoop tests the CircularArrayLoop function.
func TestCircularArrayLoop(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("CircularArrayLoop(%v)", tc.nums), func(t *testing.T) {
			actual := CircularArrayLoop(tc.nums)
			if actual != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, actual)
			}
		})
	}
}

func BenchmarkCircularArrayLoop(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			CircularArrayLoop(tc.nums)
		}
	}
}
