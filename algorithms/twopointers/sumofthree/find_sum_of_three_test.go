package sumofthree

import (
	"fmt"
	"testing"
)

type testCase struct {
	numbers  []int
	target   int
	expected bool
}

// TestCases for the findSomeOfThree function
var testCases = []testCase{
	{[]int{1, -1, 0}, -1, false},
	{[]int{3, 7, 1, 2, 8, 4, 5}, 21, false},
	{[]int{3, 7, 1, 2, 8, 4, 5}, 10, true},
	{[]int{-1, 2, 1, -4, 5, -3}, -8, true},
	{[]int{-1, 2, 1, -4, 5, -3}, 0, true},
}

func TestFindSumOfThree(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("findSumOfThree(%v, %d)", tc.numbers, tc.numbers), func(t *testing.T) {
			result := findSomeOfThree(tc.numbers, tc.target)
			if result != tc.expected {
				t.Errorf("For input %v, expected %v but got %v", tc.numbers, tc.expected, result)
			}
		})
	}
}

func BenchmarkFindSumOfThree(b *testing.B) {
	if testing.Short() {
		b.Skip("Skipping benchmark in short mode.")
	}

	for _, tc := range testCases {
		for i := 0; i < b.N; i++ {
			b.Run(fmt.Sprintf("findSumOfThree(%v, %d)", tc.numbers, tc.numbers), func(b *testing.B) {
				result := findSomeOfThree(tc.numbers, tc.target)
				if result != tc.expected {
					b.Errorf("For input %v, expected %v but got %v", tc.numbers, tc.expected, result)
				}
			})
		}
	}
}
