package search2dmatrix

import (
	"fmt"
	"testing"
)

type testCase struct {
	matrix   [][]int
	target   int
	expected bool
}

var search2dMatrixIITestCases = []testCase{
	{
		matrix: [][]int{
			{-7, -1, 5, 11},
			{-6, 0, 6, 14},
			{-4, 2, 9, 20},
		},
		target:   6,
		expected: true,
	},
	{
		matrix: [][]int{
			{1, 3, 8},
			{2, 5, 9},
			{4, 7, 12},
		},
		target:   6,
		expected: false,
	},
	{
		matrix: [][]int{
			{-2, 0, 3},
			{1, 4, 7},
		},
		target:   -5,
		expected: false,
	},
	{
		matrix: [][]int{
			{-10, -3},
			{-4, 2},
			{0, 5},
		},
		target:   9,
		expected: false,
	},
	{
		matrix: [][]int{
			{-13},
		},
		target:   -13,
		expected: true,
	},
}

func TestSearch2DMatrixII(t *testing.T) {
	for _, tc := range search2dMatrixIITestCases {
		t.Run(fmt.Sprintf("search2dMatrixII(%v, %d)", tc.matrix, tc.target), func(t *testing.T) {
			actual := searchMatrix(tc.matrix, tc.target)
			if actual != tc.expected {
				t.Fatalf("Expected %v, but got %v", tc.expected, actual)
			}
		})
	}
}

func BenchmarkSearch2DMatrixII(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for b.Loop() {
		for _, tc := range search2dMatrixIITestCases {
			b.Run(fmt.Sprintf("search2dMatrixII(%v, %d)", tc.matrix, tc.target), func(t *testing.B) {
				actual := searchMatrix(tc.matrix, tc.target)
				if actual != tc.expected {
					t.Fatalf("Expected %v, but got %v", tc.expected, actual)
				}
			})
		}
	}
}
