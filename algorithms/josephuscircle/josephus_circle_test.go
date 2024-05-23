package josephuscircle

import (
	"fmt"
	"gopherland/datastructures/list"
	"gopherland/datastructures/list/circularlinked"
	"testing"
)

func TestJosephusCircle(t *testing.T) {
	type testCase[T comparable] struct {
		data     []T
		step     int
		expected list.Node[T]
	}

	testCases := []testCase[any]{
		{
			data: []any{1, 2, 3, 4},
			step: 2,
			expected: list.Node[any]{
				Data: 1,
			},
		},
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("for list of %v and step of %d, it should return %v", tc.data, tc.step, tc.expected)
		t.Run(testName, func(t *testing.T) {
			circularList := circularlinked.New[any]()
			for _, d := range tc.data {
				circularList.Append(d)
			}

			actual := josephusCircle(*circularList, tc.step)

			if !actual.Compare(&tc.expected) {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
