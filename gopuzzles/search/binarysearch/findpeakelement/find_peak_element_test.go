package findpeakelement

import "testing"

type testCase struct {
	nums     []int
	expected int
}

var testCases = []testCase{
	{
		nums:     []int{1, 2, 3, 1},
		expected: 2,
	},
	{
		nums:     []int{1, 2, 1, 3, 5, 6, 4},
		expected: 5,
	},
	{
		nums:     []int{0, 1, 2, 3, 2, 1, 0},
		expected: 3,
	},
	{
		nums:     []int{0, 10, 3, 2, 1, 0},
		expected: 1,
	},
	{
		nums:     []int{0, 10, 0},
		expected: 1,
	},
	{
		nums:     []int{0, 1, 2, 12, 22, 32, 42, 52, 62, 72, 82, 92, 102, 112, 122, 132, 133, 132, 111, 0},
		expected: 16,
	},
	{
		nums:     []int{0, 10, 5, 2},
		expected: 1,
	},
	{
		nums:     []int{0, 2, 1, 0},
		expected: 1,
	},
	{
		nums:     []int{0, 1, 0},
		expected: 1,
	},
}

func TestFindPeakElement(t *testing.T) {
	for _, tc := range testCases {
		actual := findPeakElement(tc.nums)
		if actual != tc.expected {
			t.Fatalf("findPeakElement(nums=%d)=%d expected=%d, got=", tc.nums, actual, tc.expected)
		}
	}
}

func BenchmarkFindPeakElement(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := findPeakElement(tc.nums)
			if actual != tc.expected {
				b.Fatalf("findPeakElement(nums=%d)=%d expected=%d, got=", tc.nums, actual, tc.expected)
			}
		}
	}
}
