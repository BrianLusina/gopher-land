package findduplicate

import (
	"fmt"
	"testing"
)

type testCase struct {
	numbers  []int
	expected int
}

var testCases = []testCase{
	{
		numbers:  []int{1, 3, 4, 2, 2},
		expected: 2,
	},
	{
		numbers:  []int{3, 1, 3, 4, 2},
		expected: 3,
	},
	{
		numbers:  []int{3, 4, 1, 4, 2},
		expected: 4,
	},
	{
		numbers:  []int{1, 2, 3},
		expected: -1,
	},
	{
		numbers:  []int{3, 4, 1, 4, 1},
		expected: 4,
	},
	{
		numbers:  []int{1, 3, 3, 4, 2, 5},
		expected: 3,
	},
	{
		numbers:  []int{1, 5, 3, 4, 2, 5},
		expected: 5,
	},
	{
		numbers:  []int{1, 2, 3, 4, 5, 6, 6, 7},
		expected: 6,
	},
	{
		numbers:  []int{4, 6, 7, 7, 3, 5, 2, 8, 1},
		expected: 7,
	},
	{
		numbers:  []int{9, 8, 7, 6, 2, 3, 5, 4, 1, 9},
		expected: 9,
	},
	{
		numbers:  []int{3, 4, 4, 4, 2},
		expected: 4,
	},
	{
		numbers:  []int{1, 1},
		expected: 1,
	},
	{
		numbers:  []int{1, 3, 4, 2, 2},
		expected: 2,
	},
	{
		numbers:  []int{1, 3, 6, 2, 7, 3, 5, 4},
		expected: 3,
	},
	{
		numbers:  []int{1, 2, 2},
		expected: 2,
	},
}

func TestFindDuplicateSet(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("numbers = %v", tc.numbers), func(t *testing.T) {
			actual := findDuplicateSet(tc.numbers)
			if actual != tc.expected {
				t.Fatalf("findDuplicate(%v) = %d, expected = %d", tc.numbers, actual, tc.expected)
			}
		})
	}
}

func BenchmarkFindDuplicateSet(b *testing.B) {
	for i := 0; i < b.N; i++ {

		for _, test := range testCases {
			findDuplicateSet(test.numbers)
		}
	}
}

func TestFindDuplicate(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("numbers = %v", tc.numbers), func(t *testing.T) {
			actual := findDuplicate(tc.numbers)
			if actual != tc.expected {
				t.Fatalf("findDuplicate(%v) = %d, expected = %d", tc.numbers, actual, tc.expected)
			}
		})
	}
}

func BenchmarkFindDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {

		for _, test := range testCases {
			findDuplicate(test.numbers)
		}
	}
}
