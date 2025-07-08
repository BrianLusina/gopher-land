package sortedsquaredarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	array    []int
	expected []int
}

var testCases = []testCase{
	{
		array:    []int{-4, -1, 0, 3, 10},
		expected: []int{0, 1, 9, 16, 100},
	},
	{
		array:    []int{-7, -3, 2, 3, 11},
		expected: []int{4, 9, 9, 49, 121},
	},
	{
		array:    []int{-10, -5, 0, 5, 10},
		expected: []int{0, 25, 25, 100, 100},
	},
}

func TestSortedSquaredArray(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("SortedSquaredArray(%v)", tc.array), func(t *testing.T) {
			actual := SortedSquaredArray(tc.array)
			if !assert.Equal(t, tc.expected, actual) {
				t.Errorf("expected %v but got %v", tc.expected, actual)
			}
		})
	}
}

func BenchmarkSortedSquaredArray(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for b.Loop() {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("SortedSquaredArray(%v)", tc.array), func(b *testing.B) {
				actual := SortedSquaredArray(tc.array)
				if !assert.Equal(b, tc.expected, actual) {
					b.Errorf("expected %v but got %v", tc.expected, actual)
				}
			})
		}
	}
}

func TestSortedSquaredArray2(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("SortedSquaredArray2(%v)", tc.array), func(t *testing.T) {
			actual := SortedSquaredArray2(tc.array)
			if !assert.Equal(t, tc.expected, actual) {
				t.Errorf("expected %v but got %v", tc.expected, actual)
			}
		})
	}
}

func BenchmarkSortedSquaredArray2(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for b.Loop() {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("SortedSquaredArray2(%v)", tc.array), func(b *testing.B) {
				actual := SortedSquaredArray2(tc.array)
				if !assert.Equal(b, tc.expected, actual) {
					b.Errorf("expected %v but got %v", tc.expected, actual)
				}
			})
		}
	}
}
