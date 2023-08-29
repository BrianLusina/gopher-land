package trappedrain

import (
	"fmt"
	"testing"
)

type testCase struct {
	height   []int
	expected int
}

var testCases = []testCase{
	{
		height:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		expected: 6,
	},
	{
		height:   []int{1, 2},
		expected: 0,
	},
	{
		height:   []int{4, 2, 0, 3, 2, 5},
		expected: 9,
	},
}

func TestTrappedRainWater(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("trap(%v) should return %d", tc.height, tc.expected), func(t *testing.T) {
			actual := trap(tc.height)
			if actual != tc.expected {
				t.Fatalf("got %d, expected %d", actual, tc.expected)
			}
		})
	}
}

func BenchmarkTrappedRainWater(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("trap(%v) should return %d", tc.height, tc.expected), func(t *testing.B) {
				actual := trap(tc.height)
				if actual != tc.expected {
					t.Fatalf("got %d, expected %d", actual, tc.expected)
				}
			})
		}
	}
}
