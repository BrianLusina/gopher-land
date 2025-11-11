package settozero

import (
	"fmt"
	"gopherland/pkg/utils"
	"testing"
)

type testCase struct {
	matrix   [][]int
	expected [][]int
}

var testCases = []testCase{
	{
		matrix:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 0, 9}},
		expected: [][]int{{1, 0, 3}, {4, 0, 6}, {0, 0, 0}},
	},
	{
		matrix:   [][]int{{1, 2, 3, 4}, {4, 5, 6, 7}, {8, 9, 4, 6}},
		expected: [][]int{{1, 2, 3, 4}, {4, 5, 6, 7}, {8, 9, 4, 6}},
	},
	{
		matrix:   [][]int{{1, 1, 0, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 0, 1, 1, 1}, {1, 1, 1, 1, 1}},
		expected: [][]int{{0, 0, 0, 0, 0}, {1, 0, 0, 1, 1}, {1, 0, 0, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 1, 1}},
	},
	{
		matrix:   [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
		expected: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
	},
	{
		matrix:   [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
		expected: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
	},
	{
		matrix:   [][]int{{3, 5, 2, 0}, {1, 0, 4, 6}, {7, 3, 2, 4}},
		expected: [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {7, 0, 2, 0}},
	},
}

func TestSetMatrixZero(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Matrix=%v", tc.matrix), func(t *testing.T) {
			actual := setMatrixZeros(tc.matrix)
			expected := tc.expected
			if !utils.EqualMatrices(actual, expected) {
				t.Errorf("setMatrixZeros(%v) = %v; expected %v", tc.matrix, actual, expected)
			}
		})
	}
}

func BenchmarkSetMatrixZero(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for b.Loop() {
		for _, tc := range testCases {
			setMatrixZeros(tc.matrix)
		}
	}
}
