package finddiffoftwoarrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	nums1    []int
	nums2    []int
	expected [][]int
}

var testCases = []testCase{
	{
		nums1:    []int{1, 2, 3},
		nums2:    []int{2, 4, 6},
		expected: [][]int{{1, 3}, {4, 6}},
	},
	{
		nums1:    []int{1, 2, 3, 3},
		nums2:    []int{1, 1, 2, 2},
		expected: [][]int{{3}, {}},
	},
}

func TestFindDifference(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("findDifference(%v, %v)", tc.nums1, tc.nums2), func(t *testing.T) {
			actual := findDifference(tc.nums1, tc.nums2)
			if !assert.ElementsMatch(t, actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func BenchmarkFindDifference(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("findDifference(%v, %v)", tc.nums1, tc.nums2), func(b *testing.B) {
				actual := findDifference(tc.nums1, tc.nums2)
				if !assert.ElementsMatch(b, actual, tc.expected) {
					b.Errorf("expected %v, got %v", tc.expected, actual)
				}
			})
		}
	}
}
