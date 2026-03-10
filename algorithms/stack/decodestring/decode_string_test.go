package decodestring

import (
	"fmt"
	"testing"
)

type testCase struct {
	encodedString string
	expected      string
}

var testCases = []testCase{
	{
		encodedString: "3[a]2[bc]",
		expected:      "aaabcbc",
	},
	{
		encodedString: "3[a2[c]]",
		expected:      "accaccacc",
	},
	{
		encodedString: "2[abc]3[cd]ef",
		expected:      "abcabccdcdcdef",
	},
}

func TestDecodeString(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("decodeString(%s)", tc.encodedString), func(t *testing.T) {
			actual := decodeString(tc.encodedString)
			if actual != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, actual)
			}
		})
	}
}

func BenchmarkDecodeString(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("decodeString(%s)", tc.encodedString), func(b *testing.B) {
				actual := decodeString(tc.encodedString)
				if actual != tc.expected {
					b.Errorf("expected %s, got %s", tc.expected, actual)
				}
			})
		}
	}
}
