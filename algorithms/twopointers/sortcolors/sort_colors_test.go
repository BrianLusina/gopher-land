package sortcolors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	colors   []int
	expected []int
}

var testCases = []testCase{
	{
		colors:   []int{1, 0, 2, 1, 2, 2},
		expected: []int{0, 1, 1, 2, 2, 2},
	},
	{
		colors:   []int{0, 1, 1, 2, 0, 2, 0, 2, 1, 2},
		expected: []int{0, 0, 0, 1, 1, 1, 2, 2, 2, 2},
	},
	{
		colors:   []int{0},
		expected: []int{0},
	},
	{
		colors:   []int{0, 1, 0},
		expected: []int{0, 0, 1},
	},
	{
		colors:   []int{1},
		expected: []int{1},
	},
	{
		colors:   []int{2, 2},
		expected: []int{2, 2},
	},
	{
		colors:   []int{2, 2},
		expected: []int{2, 2},
	},
	{
		colors:   []int{1, 1, 0, 2},
		expected: []int{0, 1, 1, 2},
	},
	{
		colors:   []int{2, 1, 1, 0, 0},
		expected: []int{0, 0, 1, 1, 2},
	},
}

func TestSortColors(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("sortColors(%v)", tc.colors), func(t *testing.T) {
			actual := sortColors(tc.colors)
			assert.ElementsMatch(t, tc.expected, actual)
		})
	}
}

func BenchmarkSortColors(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("sortColors(%v)", tc.colors), func(b *testing.B) {
				actual := sortColors(tc.colors)
				assert.ElementsMatch(b, tc.expected, actual)
			})
		}
	}
}
