package pivotindex

import "testing"

type testCase struct {
	nums     []int
	expected int
}

var testCases = []testCase{
	{
		nums:     []int{1, 7, 3, 6, 5, 6},
		expected: 3,
	},
	{
		nums:     []int{1, 2, 3},
		expected: -1,
	},
	{
		nums:     []int{2, 1, -1},
		expected: 0,
	},
}

func TestPivotIndex(t *testing.T) {
	for _, tc := range testCases {
		actual := pivotIndex(tc.nums)
		if actual != tc.expected {
			t.Errorf("pivotIndex(%v) = %d expected %d", tc.nums, actual, tc.expected)
		}
	}
}

func TestPivotIndexOne(t *testing.T) {
	for _, tc := range testCases {
		actual := pivotIndexOne(tc.nums)
		if actual != tc.expected {
			t.Errorf("pivotIndex(%v) = %d expected %d", tc.nums, actual, tc.expected)
		}
	}
}

func BenchmarkPivotIndex(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark")
	}
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			pivotIndex(testCase.nums)
		}
	}
}

func BenchmarkPivotIndexOne(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark")
	}
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			pivotIndexOne(testCase.nums)
		}
	}
}
