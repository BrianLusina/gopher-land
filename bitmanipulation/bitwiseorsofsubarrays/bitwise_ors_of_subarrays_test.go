package bitwiseorsofsubarrays

import "testing"

type testCase struct {
	input    []int
	expected int
}

var testCases = []testCase{
	{input: []int{0}, expected: 1},
	{input: []int{1, 1, 2}, expected: 3},
	{input: []int{1, 2, 4}, expected: 6},
	{input: []int{37}, expected: 1},
	{input: []int{0, 0, 0, 0, 0}, expected: 1},
	{input: []int{6, 6, 6, 6}, expected: 1},
	{input: []int{1, 2, 4, 8}, expected: 10},
	{input: []int{0, 3, 0, 5}, expected: 4},
}

func TestSubarrayBitwiseORs(t *testing.T) {
	for _, tc := range testCases {
		result := subarrayBitwiseORs(tc.input)
		if result != tc.expected {
			t.Errorf("For input %v, expected %d but got %d", tc.input, tc.expected, result)
		}
	}
}
