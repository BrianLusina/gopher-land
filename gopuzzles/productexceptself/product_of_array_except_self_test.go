package productexceptself

import (
	"reflect"
	"testing"
)

type testCase struct {
	input  []int
	output []int
}

var testCases = []testCase{
	{
		input:  []int{1, 2, 3, 4},
		output: []int{24, 12, 8, 6},
	},
	{
		input:  []int{-1, 1, 0, -3, 3},
		output: []int{0, 0, 9, 0, 0},
	},
}

func TestProductExceptSelf(t *testing.T) {
	for _, tc := range testCases {
		actual := productExceptSelf(tc.input)

		if !reflect.DeepEqual(actual, tc.output) {
			t.Fatalf("productExceptSelf(%v) = %v, expected %v", tc.input, actual, tc.output)
		}
	}
}
