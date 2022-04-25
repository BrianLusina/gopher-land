package jumpgame

import "testing"

type testCases struct {
	desc     string
	input    []int
	expected int
}

func TestJumpGameIv(t *testing.T) {
	tests := []testCases{
		{
			desc:     "should return min of 3 jumps from 100, -23, -23, 404, 100, 23, 23, 23, 3, 404",
			input:    []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404},
			expected: 3,
		},
		{
			desc:     "should return min of 0 jumps from array with 1 value",
			input:    []int{7},
			expected: 0,
		},
		{
			desc:     "should return min of 1 jump from 7,6,9,6,9,6,9,7",
			input:    []int{7, 6, 9, 6, 9, 6, 9, 7},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			actual := minJumps(tt.input)
			if actual != tt.expected {
				t.Errorf("got %d, expected %d", actual, tt.expected)
			}
		})
	}
}
