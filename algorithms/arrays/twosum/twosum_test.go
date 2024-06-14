package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	numbers  []int
	target   int
	expected [2]int
}

var testCases = []testCase{
	{
		numbers:  []int{2, 7, 11, 15},
		target:   9,
		expected: [2]int{1, 2},
	},
	{
		numbers:  []int{2, 3, 4},
		target:   6,
		expected: [2]int{1, 3},
	},
	{
		numbers:  []int{-1, 0},
		target:   -1,
		expected: [2]int{1, 2},
	},
}

func TestTwoSumII(t *testing.T) {
	for _, tc := range testCases {
		if actual := twoSumII(tc.numbers, tc.target); !assert.ElementsMatch(t, actual, tc.expected) {
			t.Fatalf("twoSumII(%v, %d)=%v, expected=%v", tc.numbers, tc.target, actual, tc.expected)
		}
	}
}

func BenchmarkTwoSumII(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			if actual := twoSumII(tc.numbers, tc.target); !assert.ElementsMatch(b, actual, tc.expected) {
				b.Fatalf("twoSumII(%v, %d)=%v, expected=%v", tc.numbers, tc.target, actual, tc.expected)
			}
		}
	}
}
