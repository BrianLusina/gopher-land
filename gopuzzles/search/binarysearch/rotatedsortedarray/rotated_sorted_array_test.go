package rotatedsortedarray

import "testing"

type testCase struct {
	nums     []int
	target   int
	expected int
}

var testCases = []testCase{
	{
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		target:   0,
		expected: 4,
	},
	{
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		target:   3,
		expected: -1,
	},
	{
		nums:     []int{1},
		target:   0,
		expected: -1,
	},
}

func TestSearchRotatedSortedArray(t *testing.T) {
	for _, tc := range testCases {
		actual := search(tc.nums, tc.target)
		if actual != tc.expected {
			t.Fatalf("search(%v, %d) = %d, expected %d", tc.nums, tc.target, actual, tc.expected)
		}
	}
}
