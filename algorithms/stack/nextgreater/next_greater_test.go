package nextgreater

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	nums     []int
	expected []int
}

var testCases = []testCase{
	{
		nums:     []int{4, 5, 2, 10},
		expected: []int{5, 10, 10, -1},
	},
	{
		nums:     []int{3, 2, 1},
		expected: []int{-1, -1, -1},
	},
	{
		nums:     []int{4, 5, 2, 25},
		expected: []int{5, 25, 25, -1},
	},
	{
		nums:     []int{13, 7, 6, 12},
		expected: []int{-1, 12, 12, -1},
	},
	{
		nums:     []int{1},
		expected: []int{-1},
	},
	{
		nums:     []int{34, 35, 27, 42, 5, 28, 39, 20, 28},
		expected: []int{35, 42, 42, -1, 28, 39, -1, 28, -1},
	},
	{
		nums:     []int{11, 13, 21, 3},
		expected: []int{13, 21, -1, -1},
	},
	{
		nums:     []int{6, 8, 0, 1, 3},
		expected: []int{8, -1, 1, 3, -1},
	},
	{
		nums:     []int{3, 8, 4, 1, 2, 6, 7, 2},
		expected: []int{8, -1, 6, 2, 6, 7, -1, -1},
	},
	{
		nums:     []int{1, 3, 2, 4},
		expected: []int{3, 4, 4, -1},
	},
	{
		nums:     []int{7, 8, 1, 4},
		expected: []int{8, -1, 4, -1},
	},
}

func TestNextGreaterElement(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("nextGreaterElement(%v)", tc.nums), func(t *testing.T) {
			actual := nextGreaterElement(tc.nums)
			if !assert.ElementsMatch(t, actual, tc.expected) {
				t.Fatalf("expected = %v, got = %v", tc.expected, actual)
			}
		})
	}
}

func BenchmarkNextGreaterElement(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("nextGreaterElement(%v)", tc.nums), func(b *testing.B) {
				actual := nextGreaterElement(tc.nums)
				if !assert.ElementsMatch(b, actual, tc.expected) {
					b.Fatalf("expected = %v, got = %v", tc.expected, actual)
				}
			})
		}
	}
}
