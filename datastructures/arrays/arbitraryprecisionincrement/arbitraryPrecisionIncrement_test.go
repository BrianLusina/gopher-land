package arbitraryprecisionincrement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	a        []int
	expected []int
}

var testCases = []testCase{
	{
		a:        []int{1, 4, 9},
		expected: []int{1, 5, 0},
	},
	{
		a:        []int{9, 9, 9},
		expected: []int{1, 0, 0, 0},
	},
}

func TestArbitraryPrecisionIncrement(t *testing.T) {
	for _, tc := range testCases {
		testName := fmt.Sprintf("should return %v for input of %v", tc.expected, tc.a)
		t.Run(testName, func(t *testing.T) {
			actual := arbitraryPrecisionIncrement(tc.a)
			assert.ElementsMatch(t, actual, tc.expected)
		})
	}
}

func BenchmarkArbitraryPrecisionIncrement(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("should return %v for input of %v", tc.expected, tc.a)
		b.Run(testName, func(b *testing.B) {
			actual := arbitraryPrecisionIncrement(tc.a)
			assert.ElementsMatch(b, actual, tc.expected)
		})
	}
}
