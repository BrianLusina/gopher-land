package arrayadvance

import (
	"fmt"
	"testing"
)

type testCase struct {
	a        []int
	expected bool
}

var testCases = []testCase{
	{
		a:        []int{3, 3, 1, 0, 2, 0, 1},
		expected: true,
	},
	{
		a:        []int{3, 2, 0, 0, 2, 0, 1},
		expected: false,
	},
	{
		a:        []int{1, 4, 1, 1, 0, 2, 3},
		expected: true,
	},
}

func TestArrayAdvance(t *testing.T) {
	for _, tc := range testCases {
		testName := fmt.Sprintf("should return %v for %v", tc.expected, tc.a)
		t.Run(testName, func(t *testing.T) {
			actual := arrayAdvance(tc.a)
			if tc.expected != actual {
				t.Fail()
			}
		})
	}
}

func BenchmarkArrayAdvance(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("should return %v for %v", tc.expected, tc.a)
		b.Run(testName, func(b *testing.B) {
			actual := arrayAdvance(tc.a)
			if tc.expected != actual {
				b.Fail()
			}
		})
	}
}
