package singlenumber

import "testing"

type testCase struct {
	nums     []int
	expected int
}

var testCases = []testCase{
	{
		nums:     []int{2, 2, 1},
		expected: 1,
	},
	{
		nums:     []int{4, 1, 2, 1, 2},
		expected: 4,
	},
	{
		nums:     []int{1},
		expected: 1,
	},
}

func TestSingleNumber(t *testing.T) {
	for _, tc := range testCases {
		actual := singleNumber(tc.nums)
		if actual != tc.expected {
			t.Fatalf("singleNumber(nums=%v) = %d, expected: %d", tc.nums, actual, tc.expected)
		}
	}
}

func TestSingleNumberMath(t *testing.T) {
	for _, tc := range testCases {
		actual := singleNumberMath(tc.nums)
		if actual != tc.expected {
			t.Fatalf("singleNumber(nums=%v) = %d, expected: %d", tc.nums, actual, tc.expected)
		}
	}
}

func BenchmarkSingleNumber(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := singleNumber(tc.nums)
			if actual != tc.expected {
				b.Fatalf("singleNumber(nums=%v) = %d, expected: %d", tc.nums, actual, tc.expected)
			}
		}
	}
}

func BenchmarkSingleNumberMath(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := singleNumberMath(tc.nums)
			if actual != tc.expected {
				b.Fatalf("singleNumber(nums=%v) = %d, expected: %d", tc.nums, actual, tc.expected)
			}
		}
	}
}
