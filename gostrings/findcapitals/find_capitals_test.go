package findcapitals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	word     string
	expected []int
}

var testCases = []testCase{
	{
		word:     "CodEWaRs",
		expected: []int{0, 3, 4, 6},
	},
}

func TestFindCapitals(t *testing.T) {
	for _, tc := range testCases {
		actual := findCapitals(tc.word)
		if !assert.Equal(t, tc.expected, actual) {
			t.Fatalf("findCapitals(%s) = %v is not equal to expected %v", tc.word, actual, tc.expected)
		}
	}
}

func BenchmarkFindCapitals(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark")
	}

	for _, tc := range testCases {
		b.Run(tc.word, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				findCapitals(tc.word)
			}
		})
	}
}
