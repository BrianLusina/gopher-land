package maxsubarray

import "testing"

type testCase struct {
	input    []int
	expected int
}

var testCases = []testCase{
	{
		input:    []int{},
		expected: 0,
	},
	{
		input:    []int{1},
		expected: 1,
	},
	{
		input:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		expected: 6,
	},
	{
		input:    []int{5, 4, -1, 7, 8},
		expected: 23,
	},
	{
		input:    []int{5, 4, -1, 7, 8},
		expected: 23,
	},
	{
		input:    []int{1, 2, 3, 4, -10},
		expected: 10,
	},
	{
		input:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		expected: 6,
	},
	{
		input:    []int{-500},
		expected: -500,
	},
}

func TestMaxSubarray(t *testing.T) {
	for _, tc := range testCases {
		actual := MaxSubarray(tc.input)
		if actual != tc.expected {
			t.Fatalf("maxsubarry(%v) = %d is not equal to %d", tc.input, actual, tc.expected)
		}
	}
}

func BenchmarkMaxSubArray(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			MaxSubarray(tt.input)
		}
	}
}
