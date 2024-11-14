package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	type testCase struct {
		input    string
		expected string
	}

	testCases := []testCase{
		{"HnDMao", "oaMDnH"},
		{"ClNNxX", "XxNNlC"},
		{"iRvxxH", "HxxvRi"},
		{"bqTVvA", "AvVTqb"},
		{"wvSyRu", "uRySvw"},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("should reverse string %s to %s", testCase.input, testCase.expected), func(t *testing.T) {
			input := testCase.input
			expected := testCase.expected
			actual := ReverseString(input)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestRemoveNewLineSuffice(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected string
	}{
		"empty": {
			input:    "",
			expected: "",
		},
		"ending with \r\n": {
			input:    "a\r\n",
			expected: "a",
		},
		"ending with \n": {
			input:    "a\n",
			expected: "a",
		},
		"ending with multiple \n": {
			input:    "a\n\n\n",
			expected: "a",
		},
		"ending without new line": {
			input:    "a",
			expected: "a",
		},
	}

	for name, testCase := range testCases {
		test := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := RemoveNewLineSuffixes(testCase.input)
			if actual != test.expected {
				t.Errorf("got %s, expected: %s", actual, test.expected)
			}
		})
	}
}
