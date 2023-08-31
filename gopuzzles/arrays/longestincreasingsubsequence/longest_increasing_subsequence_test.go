package longestincreasingsubsequence

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
		nums:     []int{1, 2, 1, 5},
		expected: 3,
	},
	{
		nums:     []int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15},
		expected: 6,
	},
	{
		nums:     []int{3, 10, 2, 1, 20},
		expected: 3,
	},
	{
		nums:     []int{3, 2},
		expected: 1,
	},
	{
		nums:     []int{50, 3, 10, 7, 40, 80},
		expected: 4,
	},
	{
		nums:     []int{10, 9, 2, 5, 3, 7, 101, 18},
		expected: 4,
	},
	{
		nums:     []int{0, 1, 0, 3, 2, 3},
		expected: 4,
	},
	{
		nums:     []int{7, 7, 7, 7, 7, 7, 7},
		expected: 1,
	},
}

func TestLongestIncreasingSubsequence(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("longestIncreasingSubsequence(%v)", tc.nums), func(t *testing.T) {
			actual := longestIncreasingSubsequence(tc.nums)
			if actual != tc.expected {
				t.Fatalf("expected %d got %d", tc.expected, actual)
			}
		})
	}
}

func BenchmarkLongestIncreasingSubsequence(b *testing.B) {
	if testing.Short() {
		b.Skipf("Skipping benchmark tests")
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("longestIncreasingSubsequence(%v)", tc.nums), func(b *testing.B) {
				actual := longestIncreasingSubsequence(tc.nums)
				if actual != tc.expected {
					b.Fatalf("expected %d got %d", tc.expected, actual)
				}
			})
		}
	}
}
