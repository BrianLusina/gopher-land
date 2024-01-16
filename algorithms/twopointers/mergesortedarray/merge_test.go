package mergesortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	nums1    []int
	m        int
	nums2    []int
	n        int
	expected []int
}

var testCases = []testCase{
	{
		nums1:    []int{},
		m:        0,
		nums2:    []int{},
		n:        0,
		expected: []int{},
	},
	{
		nums1:    []int{1, 2, 3, 0, 0, 0},
		m:        3,
		nums2:    []int{2, 5, 6},
		n:        3,
		expected: []int{1, 2, 2, 3, 5, 6},
	},
	{
		nums1:    []int{0},
		m:        0,
		nums2:    []int{1},
		n:        1,
		expected: []int{1},
	},
}

func TestMerge(t *testing.T) {
	for _, tc := range testCases {
		merge(tc.nums1, tc.m, tc.nums2, tc.n)
		assert.ElementsMatch(t, tc.expected, tc.nums1)
	}
}

func BenchmarkMerge(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			merge(tc.nums1, tc.m, tc.nums2, tc.n)
		}
	}
}
