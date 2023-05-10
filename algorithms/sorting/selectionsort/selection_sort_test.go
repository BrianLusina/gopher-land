package selectionsort

import (
	"fmt"
	"gopherland/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase[T types.Comparable] struct {
	input    []T
	expected []T
}

var testCaseInts = []testCase[int]{
	{
		input:    []int{4, 2, 3, 1, 7},
		expected: []int{1, 2, 3, 4, 7},
	},
	{
		input:    []int{5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5},
	},
}

func TestSelectionSortInts(t *testing.T) {
	for _, tc := range testCaseInts {
		actual := SelectionSort(tc.input)

		assert.Equal(t, tc.expected, actual, fmt.Sprintf("SelectionSort(%v) = %v, expected = %v", tc.input, actual, tc.expected))
	}
}
func BenchmarkSelectionSortInts(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCaseInts {
			SelectionSort(tc.input)
		}
	}
}

var testCaseFloats = []testCase[float64]{
	{
		input:    []float64{5.0, 4.3, 3.3, 2.2, 1.1},
		expected: []float64{1.1, 2.2, 3.3, 4.3, 5.0},
	},
}

func TestSelectionSortFloats(t *testing.T) {
	for _, tc := range testCaseFloats {
		actual := SelectionSort(tc.input)

		assert.Equal(t, tc.expected, actual, fmt.Sprintf("SelectionSort(%v) = %v, expected = %v", tc.input, actual, tc.expected))
	}
}

func BenchmarkSelectionSortFloats(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCaseFloats {
			SelectionSort(tc.input)
		}
	}
}

var testCaseStrings = []testCase[string]{
	{
		input:    []string{"a", "c", "d", "f", "e"},
		expected: []string{"a", "c", "d", "e", "f"},
	},
	{
		input:    []string{"b", "a", "d", "f", "e"},
		expected: []string{"a", "b", "d", "e", "f"},
	},
}

func TestSelectionSortStrings(t *testing.T) {
	for _, tc := range testCaseStrings {
		actual := SelectionSort(tc.input)

		assert.Equal(t, tc.expected, actual, fmt.Sprintf("SelectionSort(%v) = %v, expected = %v", tc.input, actual, tc.expected))
	}
}

func BenchmarkSelectionSortStrings(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCaseStrings {
			SelectionSort(tc.input)
		}
	}
}
