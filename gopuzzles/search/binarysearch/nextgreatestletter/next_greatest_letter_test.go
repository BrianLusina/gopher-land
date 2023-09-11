package nextgreatestletter

import (
	"fmt"
	"testing"
)

type testCase struct {
	letters  []byte
	target   byte
	expected byte
}

var testCases = []testCase{
	{
		letters:  []byte{'c', 'f', 'j'},
		target:   'a',
		expected: 'c',
	},
	{
		letters:  []byte{'c', 'f', 'j'},
		target:   'c',
		expected: 'f',
	},
	{
		letters:  []byte{'x', 'x', 'y', 'y'},
		target:   'z',
		expected: 'x',
	},
}

func TestNextGreatestLetter(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("nextGreatestLetter(%v, %b)", tc.letters, tc.target), func(t *testing.T) {
			actual := nextGreatestLetter(tc.letters, tc.target)
			if actual != tc.expected {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func BenchmarkNextGreatestLetter(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("NextGreatestLetter(%v, %b)", tc.letters, tc.target), func(b *testing.B) {
				actual := nextGreatestLetter(tc.letters, tc.target)
				if actual != tc.expected {
					b.Fatalf("expected %v, got %v", tc.expected, actual)
				}
			})
		}
	}
}
