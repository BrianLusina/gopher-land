package selectionsort

import (
	"fmt"
	"gopherland/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseNumber[T types.Number] struct {
	numbers  []T
	expected []T
}

var testCaseInts = []testCaseNumber[int]{
	{
		numbers:  []int{4, 2, 3, 1, 7},
		expected: []int{1, 2, 3, 4, 7},
	},
	{
		numbers:  []int{5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5},
	},
}

func TestSelectionSortInts(t *testing.T) {
	for _, tc := range testCaseInts {
		actual := SelectionSort(tc.numbers)

		assert.Equal(t, tc.expected, actual, fmt.Sprintf("SelectionSort(%v) = %v, expected = %v", tc.numbers, actual, tc.expected))
	}
}
func BenchmarkSelectionSortInts(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCaseInts {
			SelectionSort(tc.numbers)
		}
	}
}

var testCaseFloats = []testCaseNumber[float64]{
	{
		numbers:  []float64{5.0, 4.3, 3.3, 2.2, 1.1},
		expected: []float64{1.1, 2.2, 3.3, 4.3, 5.0},
	},
}

func TestSelectionSortFloats(t *testing.T) {
	for _, tc := range testCaseFloats {
		actual := SelectionSort(tc.numbers)

		assert.Equal(t, tc.expected, actual, fmt.Sprintf("SelectionSort(%v) = %v, expected = %v", tc.numbers, actual, tc.expected))
	}
}

func BenchmarkSelectionSortFloats(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCaseFloats {
			SelectionSort(tc.numbers)
		}
	}
}
