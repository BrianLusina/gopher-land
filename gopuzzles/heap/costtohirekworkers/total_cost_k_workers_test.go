package costtohirekworkers

import (
	"fmt"
	"testing"
)

type testCase struct {
	costs      []int
	k          int
	candidates int
	expected   int64
}

var testCases = []testCase{
	{
		costs:      []int{17, 12, 10, 2, 7, 2, 11, 20, 8},
		k:          3,
		candidates: 4,
		expected:   11,
	},
	{
		costs:      []int{1, 2, 4, 1},
		k:          3,
		candidates: 3,
		expected:   4,
	},
}

func TestTotalCostKWorkers(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("totalCost(costs=%v, k=%d, candidates=%d)", tc.costs, tc.k, tc.candidates), func(t *testing.T) {
			actual := totalCost(tc.costs, tc.k, tc.candidates)
			if actual != tc.expected {
				t.Fatalf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func BenchmarkTotalCostKWorkers(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := totalCost(tc.costs, tc.k, tc.candidates)
			if actual != tc.expected {
				b.Fatalf("expected %d, got %d", tc.expected, actual)
			}
		}
	}
}
