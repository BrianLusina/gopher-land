package tasksscheduler

import (
	"fmt"
	"testing"
)

type testCase struct {
	tasks    []string
	n        int
	expected int
}

var testCases = []testCase{
	{
		tasks:    []string{"A", "A", "B", "B"},
		n:        2,
		expected: 5,
	},
	{
		tasks:    []string{"A", "B", "A", "A", "B", "C"},
		n:        3,
		expected: 9,
	},
	{
		tasks:    []string{"A", "C", "B", "A"},
		n:        0,
		expected: 4,
	},
	{
		tasks:    []string{"A", "A", "A", "B", "B", "C", "C"},
		n:        1,
		expected: 7,
	},
	{
		tasks:    []string{"S", "I", "V", "U", "W", "D", "U", "X"},
		n:        0,
		expected: 8,
	},
	{
		tasks:    []string{"A", "K", "X", "M", "W", "D", "X", "B", "D", "C", "O", "Z", "D", "E", "Q"},
		n:        3,
		expected: 15,
	},
	{
		tasks:    []string{"A", "B", "C", "O", "Q", "C", "Z", "O", "X", "C", "W", "Q", "Z", "B", "M", "N", "R", "L", "C", "J"},
		n:        10,
		expected: 34,
	},
}

func TestLeastIntervalsMathematical(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("leastIntervalsMathematical(tasks=%v,n=%d)", tc.tasks, tc.n), func(t *testing.T) {
			result := leastIntervalsMathematical(tc.tasks, tc.n)
			if result != tc.expected {
				t.Errorf("For tasks %v and n=%d, expected %d but got %d", tc.tasks, tc.n, tc.expected, result)
			}
		})
	}
}

func BenchmarkLeastIntervalsMathematical(b *testing.B) {
	for b.Loop() {
		for _, tc := range testCases {
			leastIntervalsMathematical(tc.tasks, tc.n)
		}
	}
}
