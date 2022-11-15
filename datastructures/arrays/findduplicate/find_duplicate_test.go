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
