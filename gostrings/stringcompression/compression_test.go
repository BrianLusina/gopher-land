package stringcompression

import (
	"fmt"
	"testing"
)

type testCase struct {
	chars    []byte
	expected int
}

var testCases = []testCase{
	{
		chars:    []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
		expected: 6,
	},
	{
		chars:    []byte{'a'},
		expected: 1,
	},
	{
		chars:    []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
		expected: 4,
	},
}

func TestStringCompression(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("compress(%v)", tc.chars), func(t *testing.T) {
			actual := compress(tc.chars)
			if actual != tc.expected {
				t.Fatalf(fmt.Sprintf("got %d, expected %d", actual, tc.expected))
			}
		})
	}
}

func BenchmarkStringCompression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("compress(%v)", tc.chars), func(b *testing.B) {
				actual := compress(tc.chars)
				if actual != tc.expected {
					b.Fatalf(fmt.Sprintf("got %d, expected %d", actual, tc.expected))
				}
			})
		}
	}
}
