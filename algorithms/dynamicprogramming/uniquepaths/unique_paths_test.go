package uniquepaths

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	m        int
	n        int
	expected int
}

var testCases = []testCase{
	{
		m:        3,
		n:        7,
		expected: 28,
	},
	{
		m:        3,
		n:        2,
		expected: 3,
	},
	{
		m:        51,
		n:        9,
		expected: 1916797311,
	},
}

func TestUniquePathsMath(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("uniquePaths(m=%d, n=%d)", tc.m, tc.n), func(t *testing.T) {
			actual := uniquePathsMath(tc.m, tc.n)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkUniquePathsMath(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := uniquePathsMath(tc.m, tc.n)
			assert.Equal(b, tc.expected, actual)
		}
	}
}

func TestUniquePathsTopDown(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("uniquePaths(m=%d, n=%d)", tc.m, tc.n), func(t *testing.T) {
			actual := uniquePathsTopDown(tc.m, tc.n)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkUniquePathsTopDown(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := uniquePathsTopDown(tc.m, tc.n)
			assert.Equal(b, tc.expected, actual)
		}
	}
}

func TestUniquePathsBottomUp(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("uniquePaths(m=%d, n=%d)", tc.m, tc.n), func(t *testing.T) {
			actual := uniquePathsBottomUp(tc.m, tc.n)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkUniquePathsBottomUp(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := uniquePathsBottomUp(tc.m, tc.n)
			assert.Equal(b, tc.expected, actual)
		}
	}
}
