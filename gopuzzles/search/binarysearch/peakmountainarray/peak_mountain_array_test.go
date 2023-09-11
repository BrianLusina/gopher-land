package peakmountainarray

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
		numbers:  []int{0, 1, 2, 3, 2, 1, 0},
		expected: 3,
	},
	{
		numbers:  []int{0, 10, 3, 2, 1, 0},
		expected: 1,
	},
	{
		numbers:  []int{0, 10, 0},
		expected: 1,
	},
	{
		numbers:  []int{0, 1, 2, 12, 22, 32, 42, 52, 62, 72, 82, 92, 102, 112, 122, 132, 133, 132, 111, 0},
		expected: 16,
	},
	{
		numbers:  []int{0, 10, 5, 2},
		expected: 1,
	},
	{
		numbers:  []int{0, 10, 5, 2},
		expected: 1,
	},
	{
		numbers:  []int{0, 2, 1, 0},
		expected: 1,
	},
	{
		numbers:  []int{0, 1, 0},
		expected: 1,
	},
}

func TestPeakIndexInMountainArray(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("peakIndexInMountainArray(%d)", tc.numbers), func(t *testing.T) {
			actual := peakIndexInMountainArray(tc.numbers)

			if actual != tc.expected {
				t.Fatalf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func BenchmarkPeakIndexInMountainArray(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("peakIndexInMountainArray(%d)", tc.numbers), func(b *testing.B) {
				actual := peakIndexInMountainArray(tc.numbers)

				if actual != tc.expected {
					b.Fatalf("expected %d, got %d", tc.expected, actual)
				}
			})
		}
	}
}
