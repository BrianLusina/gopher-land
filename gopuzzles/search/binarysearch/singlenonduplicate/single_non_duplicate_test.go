package singlenonduplicate

import (
	"fmt"
	"testing"
)

type testCase struct {
	nums     []int
	expected int
}

var testCases = []testCase{
	{
		nums:     []int{1, 1, 2, 3, 3, 4, 4, 8, 8},
		expected: 2,
	},
	{
		nums:     []int{3, 3, 7, 7, 10, 11, 11},
		expected: 10,
	},
}

func TestSingleNonDuplicate(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("singleNonDuplicate(%v)", tc.nums), func(t *testing.T) {
			actual := singleNonDuplicate(tc.nums)
			if actual != tc.expected {
				t.Fatalf("expected: %d, got: %d", tc.expected, actual)
			}
		})
	}
}

func BenchmarkSingleNonDuplicate(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("singleNonDuplicate(%v)", tc.nums), func(b *testing.B) {
				actual := singleNonDuplicate(tc.nums)
				if actual != tc.expected {
					b.Fatalf("expected: %d, got: %d", tc.expected, actual)
				}
			})
		}
	}
}
