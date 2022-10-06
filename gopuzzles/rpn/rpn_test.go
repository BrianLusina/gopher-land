package rpn

import "testing"

type tests struct {
	tokens   []string
	expected int
}

var testCases []tests = []tests{
	{
		tokens:   []string{"2", "1", "+", "3", "*"},
		expected: 9,
	},
	{
		tokens:   []string{"4", "13", "5", "/", "+"},
		expected: 6,
	},
	{
		tokens:   []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
		expected: 22,
	},
}

func TestReversePolishNotation(t *testing.T) {
	for _, tc := range testCases {
		actual := evalRPN(tc.tokens)
		if actual != tc.expected {
			t.Fatalf("evalRPN(%v) = %d, expected=%d", tc.tokens, actual, tc.expected)
		}
		t.Logf("Pass: evalRPN(%v)", tc.tokens)
	}
}
