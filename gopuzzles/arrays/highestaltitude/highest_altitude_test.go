package highestaltitude

import (
	"fmt"
	"testing"
)

type testCase struct {
	gain     []int
	expected int
}

var testCases = []testCase{
	{
		gain:     []int{-5, 1, 5, 0, -7},
		expected: 1,
	},
	{
		gain:     []int{-4, -3, -2, -1, 4, 3, 2},
		expected: 0,
	},
}

func TestHighestAltitude(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("largestAltitude(%v)", tc.gain), func(t *testing.T) {
			actual := largestAltitude(tc.gain)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func BenchmarkHighestAltitude(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("largestAltitude(%v)", tc.gain), func(b *testing.B) {
				actual := largestAltitude(tc.gain)
				if actual != tc.expected {
					b.Errorf("expected %d, got %d", tc.expected, actual)
				}
			})
		}
	}
}
