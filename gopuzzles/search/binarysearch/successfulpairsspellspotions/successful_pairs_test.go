package successfulpairsspellspotions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	spells   []int
	potions  []int
	success  int64
	expected []int
}

var testCases = []testCase{
	{
		spells:   []int{5, 1, 3},
		potions:  []int{1, 2, 3, 4, 5},
		success:  7,
		expected: []int{4, 0, 3},
	},
	{
		spells:   []int{3, 1, 2},
		potions:  []int{8, 5, 8},
		success:  16,
		expected: []int{2, 0, 2},
	},
}

func TestSuccessfulPairs(t *testing.T) {
	for _, tc := range testCases {
		actual := successfulPairs(tc.spells, tc.potions, tc.success)
		assert.ElementsMatch(t, tc.expected, actual)
	}
}

func BenchmarkSuccessfulPairs(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := successfulPairs(tc.spells, tc.potions, tc.success)
			assert.ElementsMatch(b, tc.expected, actual)
		}
	}
}
