package longestsubarrayofones

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
		nums:     []int{1, 1, 0, 1},
		expected: 3,
	},
	{
		nums:     []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
		expected: 5,
	},
	{
		nums:     []int{1, 1, 1},
		expected: 2,
	},
}

func TestLongestSubArray(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("longestSubarray(nums=%v)", tc.nums), func(t *testing.T) {
			actual := longestSubarray(tc.nums)
			if actual != tc.expected {
				t.Errorf("expected %d got %d", tc.expected, actual)
			}
		})
	}
}

func BenchmarkLongestSubArray(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("longestSubarray(nums=%v)", tc.nums), func(b *testing.B) {
				actual := longestSubarray(tc.nums)
				if actual != tc.expected {
					b.Errorf("expected %d got %d", tc.expected, actual)
				}
			})
		}
	}
}
