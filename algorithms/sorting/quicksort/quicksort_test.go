package quicksort

import (
	"gopherland/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase[T types.Number] struct {
	list     []T
	expected []T
}

var intSliceTests = []testCase[int]{
	{
		list:     []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8},
		expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
}

var float64SliceTests = []testCase[float64]{
	{
		list:     []float64{9.9, 4.4, 3.4, 6.2, 1.4, 2.1, 10.5, 5.1, 7.1, 8.2},
		expected: []float64{1.4, 2.1, 3.4, 4.4, 5.1, 6.2, 7.1, 8.2, 9.9, 10.5},
	},
}

func TestQuickSortInts(t *testing.T) {
	for _, test := range intSliceTests {
		actual := Quicksort(test.list)
		result := assert.Equal(t, test.expected, actual)

		if !result {
			t.Errorf("FAIL: \nQuicksort(%v)\nreturned %v, expected %v.", test.list, actual, test.expected)
		}
	}
}

func TestQuickSortFloats(t *testing.T) {
	for _, test := range float64SliceTests {
		actual := Quicksort(test.list)
		result := assert.Equal(t, test.expected, actual)

		if !result {
			t.Errorf("FAIL: \nQuicksort(%v)\nreturned %v, expected %v.", test.list, actual, test.expected)
		}
	}
}

func BenchmarkQuickSortInts(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range intSliceTests {
			Quicksort(test.list)
		}
	}
}

func BenchmarkQuickSortFloats(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range float64SliceTests {
			Quicksort(test.list)
		}
	}
}
