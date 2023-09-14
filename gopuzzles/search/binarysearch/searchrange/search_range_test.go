package searchrange

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	nums     []int
	target   int
	expected []int
}

var testCases = []testCase{
	{
		nums:     []int{5, 7, 7, 8, 8, 10},
		target:   8,
		expected: []int{3, 4},
	},
	{
		nums:     []int{5, 7, 7, 8, 8, 10},
		target:   6,
		expected: []int{-1, -1},
	},
	{
		nums:     []int{},
		target:   0,
		expected: []int{-1, -1},
	},
	{
		nums:     []int{1},
		target:   1,
		expected: []int{0, 0},
	},
	{
		nums:     []int{1, 3},
		target:   1,
		expected: []int{0, 0},
	},
	{
		nums:     []int{1, 4},
		target:   4,
		expected: []int{1, 1},
	},
	{
		nums:     []int{1, 2, 3},
		target:   2,
		expected: []int{1, 1},
	},
	{
		nums:     []int{3, 3, 3},
		target:   3,
		expected: []int{0, 2},
	},
}

func TestSearchRange(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("searchRange(nums=%v, target=%d)", tc.nums, tc.target), func(t *testing.T) {
			actual := searchRange(tc.nums, tc.target)

			if !assert.ElementsMatch(t, actual, tc.expected) {
				t.Fatalf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func BenchmarkSearchRange(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("searchRange(nums=%v, target=%d)", tc.nums, tc.target), func(b *testing.B) {
				actual := searchRange(tc.nums, tc.target)

				if !assert.ElementsMatch(b, actual, tc.expected) {
					b.Fatalf("expected %d, got %d", tc.expected, actual)
				}
			})
		}
	}
}
