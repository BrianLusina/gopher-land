package slowestkey

import "testing"

type testCase struct {
	releaseTimes []int
	keysPressed  string
	expected     byte
}

var tests = []testCase{
	{
		releaseTimes: []int{9, 29, 49, 50},
		keysPressed:  "cbcd",
		expected:     'c',
	},
	{
		releaseTimes: []int{12, 23, 36, 46, 62},
		keysPressed:  "spuda",
		expected:     'a',
	},
	{
		releaseTimes: []int{1, 2},
		keysPressed:  "ba",
		expected:     'b',
	},
}

func TestSlowestKey(t *testing.T) {
	for _, test := range tests {
		t.Run(test.keysPressed, func(t *testing.T) {
			actual := slowestKey(test.releaseTimes, test.keysPressed)
			if actual != test.expected {
				t.Errorf("slowestKey(%v, %s) = %v, expected %v", test.releaseTimes, test.keysPressed, string(actual), string(test.expected))
			}
		})
	}
}
