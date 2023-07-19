package equalrowandcols

import (
	"fmt"
	"testing"
)

type testCase struct {
	grid     [][]int
	expected int
}

var testCases = []testCase{
	{
		grid:     [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}},
		expected: 1,
	},
	{
		grid:     [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}},
		expected: 3,
	},
	{
		grid:     [][]int{{11, 1}, {1, 11}},
		expected: 2,
	},
}

func TestEqualPairsHashMap(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("equalPairsHashmap(%v)", tc.grid), func(t *testing.T) {
			actual := equalPairsHashmap(tc.grid)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func BenchmarkEqualPairsHashMap(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("equalPairsHashmap(%v)", tc.grid), func(t *testing.B) {
				actual := equalPairsHashmap(tc.grid)
				if actual != tc.expected {
					t.Errorf("expected %d, got %d", tc.expected, actual)
				}
			})
		}
	}
}

func TestEqualPairsBruteForce(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("equalPairsBruteForce(%v)", tc.grid), func(t *testing.T) {
			actual := equalPairsBruteForce(tc.grid)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func BenchmarkEqualPairsBruteForce(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("equalPairsBruteForce(%v)", tc.grid), func(t *testing.B) {
				actual := equalPairsBruteForce(tc.grid)
				if actual != tc.expected {
					t.Errorf("expected %d, got %d", tc.expected, actual)
				}
			})
		}
	}
}

func TestEqualPairsTrieNode(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("equalPairsTrieNode(%v)", tc.grid), func(t *testing.T) {
			actual := equalPairsTrieNode(tc.grid)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func BenchmarkEqualPairsTrieNode(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("equalPairsTrieNode(%v)", tc.grid), func(t *testing.B) {
				actual := equalPairsTrieNode(tc.grid)
				if actual != tc.expected {
					t.Errorf("expected %d, got %d", tc.expected, actual)
				}
			})
		}
	}
}
