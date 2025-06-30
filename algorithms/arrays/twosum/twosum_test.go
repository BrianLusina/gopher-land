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

type testCaseTwoNumSum struct {
	numbers  []int
	target   int
	expected []int
}

var testCasesTwoNumSum = []testCaseTwoNumSum{
	{
		numbers:  []int{2, 7, 11, 15},
		target:   9,
		expected: []int{2, 7},
	},
	{
		numbers:  []int{2, 3, 4},
		target:   6,
		expected: []int{2, 4},
	},
	{
		numbers:  []int{-1, 0},
		target:   -1,
		expected: []int{-1, 0},
	},
	{
		numbers:  []int{3, 5, -4, 8, 11, 1, -1, 6},
		target:   10,
		expected: []int{-1, 11},
	},
	{
		numbers:  []int{4, 6},
		target:   10,
		expected: []int{4, 6},
	},
	{
		numbers:  []int{4, 6, 1},
		target:   5,
		expected: []int{4, 1},
	},
	{
		numbers:  []int{4, 6, 1, -3},
		target:   3,
		expected: []int{6, -3},
	},
	{
		numbers:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		target:   17,
		expected: []int{8, 9},
	},
	{
		numbers:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15},
		target:   18,
		expected: []int{3, 15},
	},
	{
		numbers:  []int{-7, -5, -3, -1, 0, 1, 3, 5, 7},
		target:   -5,
		expected: []int{-5, 0},
	},
	{
		numbers:  []int{-21, 301, 12, 4, 65, 56, 210, 356, 9, -47},
		target:   163,
		expected: []int{210, -47},
	},
	{
		numbers:  []int{-21, 301, 12, 4, 65, 56, 210, 356, 9, -47},
		target:   164,
		expected: []int{},
	},
	{
		numbers:  []int{3, 5, -4, 8, 11, 1, -1, 6},
		target:   15,
		expected: []int{},
	},
	{
		numbers:  []int{14},
		target:   15,
		expected: []int{},
	},
	{
		numbers:  []int{15},
		target:   15,
		expected: []int{},
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

	for b.Loop() {
		for _, tc := range testCases {
			if actual := twoSumII(tc.numbers, tc.target); !assert.ElementsMatch(b, actual, tc.expected) {
				b.Fatalf("twoSumII(%v, %d)=%v, expected=%v", tc.numbers, tc.target, actual, tc.expected)
			}
		}
	}
}

func TestTwoNumberSumI(t *testing.T) {
	for _, tc := range testCasesTwoNumSum {
		if actual := TwoNumberSum(tc.numbers, tc.target); !assert.ElementsMatch(t, actual, tc.expected) {
			t.Fatalf("TwoNumberSum(%v, %d)=%v, expected=%v", tc.numbers, tc.target, actual, tc.expected)
		}
	}
}

func BenchmarkTwoNumberSumI(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for b.Loop() {
		for _, tc := range testCasesTwoNumSum {
			if actual := TwoNumberSum(tc.numbers, tc.target); !assert.ElementsMatch(b, actual, tc.expected) {
				b.Fatalf("TwoNumberSum(%v, %d)=%v, expected=%v", tc.numbers, tc.target, actual, tc.expected)
			}
		}
	}
}

func TestTwoNumberSumII(t *testing.T) {
	for _, tc := range testCasesTwoNumSum {
		if actual := TwoNumberSumII(tc.numbers, tc.target); !assert.ElementsMatch(t, actual, tc.expected) {
			t.Fatalf("TwoNumberSumII(%v, %d)=%v, expected=%v", tc.numbers, tc.target, actual, tc.expected)
		}
	}
}

func BenchmarkTwoNumberSumII(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for b.Loop() {
		for _, tc := range testCasesTwoNumSum {
			if actual := TwoNumberSumII(tc.numbers, tc.target); !assert.ElementsMatch(b, actual, tc.expected) {
				b.Fatalf("TwoNumberSumII(%v, %d)=%v, expected=%v", tc.numbers, tc.target, actual, tc.expected)
			}
		}
	}
}
