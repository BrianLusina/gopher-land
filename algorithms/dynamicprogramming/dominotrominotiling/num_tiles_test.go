package dominotrominotiling

import (
	"fmt"
	"testing"
)

type testCase struct {
	n        int
	expected int
}

var testCases = []testCase{
	{
		n:        3,
		expected: 5,
	},
	{
		n:        1,
		expected: 1,
	},
	{
		n:        30,
		expected: 312342182,
	},
}

func TestNumTiles(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("numTilings(n=%d)", tc.n), func(t *testing.T) {
			actual := numTilings(tc.n)
			if actual != tc.expected {
				t.Fail()
			}
		})
	}
}

func Benchmark(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := numTilings(tc.n)
			if actual != tc.expected {
				b.Fail()
			}
		}
	}
}
