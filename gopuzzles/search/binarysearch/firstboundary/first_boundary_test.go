package firstboundary

import (
	"fmt"
	"testing"
)

type testCase struct {
	arr      []bool
	expected int
}

var testCases = []testCase{
	{
		arr:      []bool{false, false, true, true, true},
		expected: 2,
	},
	{
		arr:      []bool{true},
		expected: 0,
	},
	{
		arr:      []bool{false, false, false},
		expected: -1,
	},
	{
		arr:      []bool{true, true, true, true, true},
		expected: 0,
	},
	{
		arr:      []bool{false, true},
		expected: 1,
	},
	{
		arr:      []bool{false, false, false, false, false, false, false, false, true},
		expected: 8,
	},
}

func TestFindBoundary(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("findBoundary(%v)", tc.arr), func(t *testing.T) {
			actual := findBoundary(tc.arr)
			if actual != tc.expected {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func BenchmarkFindBoundary(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("findBoundary(%v)", tc.arr), func(b *testing.B) {
				actual := findBoundary(tc.arr)
				if actual != tc.expected {
					b.Fatalf("expected %v, got %v", tc.expected, actual)
				}
			})
		}
	}
}
