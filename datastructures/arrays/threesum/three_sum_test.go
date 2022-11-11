package threesum

import (
	"reflect"
	"testing"
)

type testCase struct {
	nums     []int
	expected [][]int
}

var testCases = []testCase{
	{
		nums:     []int{-1, 0, 1, 2, -1, -4},
		expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
	},
	{
		nums:     []int{0, 1, 1},
		expected: [][]int{},
	},
	{
		nums:     []int{0, 0, 0},
		expected: [][]int{{0, 0, 0}},
	},
}

func TestThreeSum(t *testing.T) {
	for _, tc := range testCases {
		actual := threeSum(tc.nums)

		if !reflect.DeepEqual(actual, tc.expected) {
			t.Fatalf("threeSum(%v) = %v , expected: %v", tc.nums, actual, tc.expected)
		}
	}
}

func BenchmarkThreeSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			threeSum(tc.nums)
		}
	}
}
