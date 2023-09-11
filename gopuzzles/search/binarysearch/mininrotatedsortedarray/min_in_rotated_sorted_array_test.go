package mininrotatedsortedarray

import "testing"

type testCase struct {
	nums     []int
	expected int
}

var testCases = []testCase{
	{
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		expected: 0,
	},
	{
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		expected: 0,
	},
	{
		nums:     []int{1},
		expected: 1,
	},
	{
		nums:     []int{3, 4, 5, 1, 2},
		expected: 1,
	},
	{
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		expected: 0,
	},
	{
		nums:     []int{11, 13, 15, 17},
		expected: 11,
	},
	{
		nums:     []int{1, 2},
		expected: 1,
	},
}

func TestFindMinInRotatedSortedArray(t *testing.T) {
	for _, tc := range testCases {
		actual := findMin(tc.nums)
		if actual != tc.expected {
			t.Fatalf("findMin(%v) = %d, expected %d", tc.nums, actual, tc.expected)
		}
	}
}
