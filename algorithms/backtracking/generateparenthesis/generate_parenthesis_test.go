package generateparenthesis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input  int
	output []string
}

var testCases []testCase = []testCase{
	{
		input:  1,
		output: []string{"()"},
	},
	{
		input:  3,
		output: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
	},
}

func TestGenerateParenthesis(t *testing.T) {
	for _, tc := range testCases {
		actual := generateParenthesis(tc.input)

		if !assert.ElementsMatch(t, actual, tc.output) {
			t.Fatalf("generateParenthesis(%d) = %v, expected = %v", tc.input, actual, tc.output)
		}

		t.Logf("Pass: generateParenthesis(%d)", tc.input)
	}
}
