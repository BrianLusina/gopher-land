package reversestring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	text     string
	expected string
}

var testCases = []testCase{
	{
		text:     "!evitacudE ot emocleW",
		expected: "Welcome to Educative!",
	},
}

func TestReverseStr(t *testing.T) {
	for _, tc := range testCases {
		actual := reverseString(tc.text)
		if actual != tc.expected {
			t.Fatalf("reverseString(%s) = %s, expected %s", tc.text, actual, tc.expected)
		}

	}
}

func BenchmarkReverseStr(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := reverseString(tc.text)
			if actual != tc.expected {
				b.Fatalf("reverseString(%s) = %s, expected %s", tc.text, actual, tc.expected)
			}
		}
	}
}

type testCaseCharArr struct {
	text     []string
	expected []string
}

var testCasesCharArray = []testCaseCharArr{
	{
		text:     []string{"a", "e", "i", "o", "u"},
		expected: []string{"u", "o", "i", "e", "a"},
	},
	{
		text:     []string{"A", "l", "e", "x"},
		expected: []string{"x", "e", "l", "A"},
	},
	{
		text:     []string{"p", "y", "t", "h", "o", "n"},
		expected: []string{"n", "o", "h", "t", "y", "p"},
	},
	{
		text:     []string{"x"},
		expected: []string{"x"},
	},
}

func TestReverseStrCharArray(t *testing.T) {
	for _, tc := range testCasesCharArray {
		sCopy := make([]string, len(tc.text))
		copy(sCopy, tc.text)
		reverseCharArray(sCopy)
		assert.Equal(t, tc.expected, sCopy)
	}
}

func BenchmarkReverseStrCharArray(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for b.Loop() {
		for _, tc := range testCasesCharArray {
			reverseCharArray(tc.text)
		}
	}
}
