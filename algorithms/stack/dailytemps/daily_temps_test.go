package dailytemps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	temperatures []int
	expected     []int
}

var testCases = []testCase{
	{
		temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
		expected:     []int{1, 1, 4, 2, 1, 1, 0, 0},
	},
	{
		temperatures: []int{30, 40, 50, 603},
		expected:     []int{1, 1, 1, 0},
	},
	{
		temperatures: []int{30, 60, 90},
		expected:     []int{1, 1, 0},
	},
}

func TestDailyTemps(t *testing.T) {
	for _, tc := range testCases {
		actual := dailyTemperatures(tc.temperatures)
		assert.ElementsMatch(t, tc.expected, actual)
	}
}

func BenchmarkDailyTemps(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := dailyTemperatures(tc.temperatures)
			assert.ElementsMatch(b, tc.expected, actual)
		}
	}
}
