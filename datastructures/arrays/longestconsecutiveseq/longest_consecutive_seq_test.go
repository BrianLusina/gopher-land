package longestconsecutiveseq

import "testing"

type testCase struct {
	nums     []int
	expected int
}

var testCases = []testCase{
	{
		nums:     []int{100, 4, 200, 1, 3, 2},
		expected: 4,
	},
	{
		nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
		expected: 9,
	},
}

func TestLongestConsecutiveSequence(t *testing.T) {
	for _, tc := range testCases {
		actual := longestConsecutive(tc.nums)
		if actual != tc.expected {
			t.Fatalf("longestConsecutive(%v) = %d, expected: %d", tc.nums, actual, tc.expected)
		}
	}
}
