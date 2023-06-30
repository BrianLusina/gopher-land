package maxaveragesubarray

import "testing"

type testCase struct {
	nums     []int
	k        int
	expected float64
}

var testCases = []testCase{
	{
		nums:     []int{1, 12, -5, -6, 50, 3},
		k:        4,
		expected: 12.75000,
	},
	{
		nums:     []int{5},
		k:        1,
		expected: 5.00000,
	},
	{
		nums:     []int{0, 1, 1, 3, 3},
		k:        4,
		expected: 2.00000,
	},
	{
		nums:     []int{4, 0, 4, 3, 3},
		k:        5,
		expected: 2.80000,
	},
}

func TestFindMaxAverageSubArray(t *testing.T) {
	for _, tc := range testCases {
		actual := findMaxAverage(tc.nums, tc.k)
		if tc.expected != actual {
			t.Fatalf("findMaxAverage(%v, %d), expected %f got %f", tc.nums, tc.k, tc.expected, actual)
		}
	}
}

func BenchmarkFindMaxAverageSubArray(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := findMaxAverage(tc.nums, tc.k)
			if tc.expected != actual {
				b.Fatalf("findMaxAverage(%v, %d), expected %f got %f", tc.nums, tc.k, tc.expected, actual)
			}
		}
	}
}
