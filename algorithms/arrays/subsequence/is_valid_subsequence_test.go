package subsequence

import (
	"fmt"
	"testing"
)

type testCase struct {
	array    []int
	sequence []int
	expected bool
}

var testCases = []testCase{
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{1, 6, -1, 10},
		expected: true,
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{5, 1, 22, 6, -1, 8, 10},
		expected: true,
	},
	{
		array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
		sequence: []int{5, 6, 1, 10, 22, 8, -1, 25},
		expected: false,
	},
	{
		array:    []int{1, 2, 3, 4},
		sequence: []int{1, 3, 4},
		expected: true,
	},
	{
		array:    []int{1, 2, 3, 4},
		sequence: []int{2, 4},
		expected: true,
	},
}

func TestIsValidSubsequence(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("IsValidSubsequence(array=%v, sequence=%v)", testCase.array, testCase.sequence), func(t *testing.T) {
			actual := IsValidSubsequence(testCase.array, testCase.sequence)
			if actual != testCase.expected {
				t.Errorf("expected %v but got %v", testCase.expected, actual)
			}
		})
	}
}

func BenchmarkIsValidSubsequence(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for b.Loop() {
		for _, testCase := range testCases {
			b.Run(fmt.Sprintf("IsValidSubsequence(array=%v, sequence=%v)", testCase.array, testCase.sequence), func(b *testing.B) {
				actual := IsValidSubsequence(testCase.array, testCase.sequence)
				if actual != testCase.expected {
					b.Errorf("expected %v but got %v", testCase.expected, actual)
				}
			})
		}
	}
}

func TestIsValidSubsequenceV2(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("IsValidSubsequenceV2(array=%v, sequence=%v)", testCase.array, testCase.sequence), func(t *testing.T) {
			actual := IsValidSubsequenceV2(testCase.array, testCase.sequence)
			if actual != testCase.expected {
				t.Errorf("expected %v but got %v", testCase.expected, actual)
			}
		})
	}
}

func BenchmarkIsValidSubsequenceV2(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for b.Loop() {
		for _, testCase := range testCases {
			b.Run(fmt.Sprintf("IsValidSubsequenceV2(array=%v, sequence=%v)", testCase.array, testCase.sequence), func(b *testing.B) {
				actual := IsValidSubsequenceV2(testCase.array, testCase.sequence)
				if actual != testCase.expected {
					b.Errorf("expected %v but got %v", testCase.expected, actual)
				}
			})
		}
	}
}
