package runningsum

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	nums     []int
	expected []int
}

var testCases = []testCase{
	{
		nums:     []int{1, 2, 3, 4},
		expected: []int{1, 3, 6, 10},
	},
	{
		nums:     []int{1, 1, 1, 1, 1},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		nums:     []int{3, 1, 2, 10, 1},
		expected: []int{3, 4, 6, 16, 17},
	},
}

func TestRunningSum(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			actual := runningSum(tc.nums)

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("runningSum(%v) = %v is not equal to %v", tc.nums, actual, tc.expected)
			}
		})
	}
}
