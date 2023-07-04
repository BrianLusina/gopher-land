package maxconsecutiveones

import (
	"fmt"
	"testing"
)

type testCase struct {
	nums     []int
	k        int
	expected int
}

var testCases = []testCase{
	{
		nums:     []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
		k:        2,
		expected: 6,
	},
	{
		nums:     []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
		k:        3,
		expected: 10,
	},
}

func TestMaxConsecutiveOnes(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("longestOnes(%v, %d)", tc.nums, tc.k), func(t *testing.T) {
			actual := longestOnes(tc.nums, tc.k)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func BenchmarkMaxConsecutiveOnes(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("longestOnes(%v, %d)", tc.nums, tc.k), func(b *testing.B) {
				actual := longestOnes(tc.nums, tc.k)
				if actual != tc.expected {
					b.Errorf("expected %d, got %d", tc.expected, actual)
				}
			})
		}
	}
}
