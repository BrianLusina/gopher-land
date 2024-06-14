package optimaltaskassignment

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	tasks    []int
	expected []Pair
}

var testCases = []testCase{
	{
		tasks:    []int{6, 3, 2, 7, 5, 5},
		expected: []Pair{{2, 7}, {3, 6}, {5, 5}},
	},
}

func TestOptimalTaskAssignment(t *testing.T) {
	for _, tc := range testCases {
		testName := fmt.Sprintf("should return %v for tasks=%v", tc.expected, tc.tasks)
		t.Run(testName, func(t *testing.T) {
			actual := optimalTaskAssignment(tc.tasks)
			assert.ElementsMatch(t, actual, tc.expected)
		})
	}
}

func BenchmarkOptimalTaskAssignment(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {

			testName := fmt.Sprintf("should return %v for tasks=%v", tc.expected, tc.tasks)
			b.Run(testName, func(b *testing.B) {
				actual := optimalTaskAssignment(tc.tasks)
				assert.ElementsMatch(b, actual, tc.expected)
			})
		}
	}
}
