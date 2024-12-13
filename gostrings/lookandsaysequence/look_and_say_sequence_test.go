package lookandsaysequence

import (
	"fmt"
	"testing"
)

type testCase struct {
	sequence string
	expected string
}

var testCases = []testCase{
	{
		sequence: "1",
		expected: "11",
	},
	{
		sequence: "11",
		expected: "21",
	},
	{
		sequence: "21",
		expected: "1211",
	},
}

func TestLookAndSaySequence(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("lookAndSaySequence(%s)", tc.sequence), func(t *testing.T) {
			actual := lookAndSaySequence(tc.sequence)
			if actual != tc.expected {
				t.Fatalf("expected %s, got: %s", tc.expected, actual)
			}
		})
	}
}

func BenchmarkLookAndSaySequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("lookAndSaySequence(%s)", tc.sequence), func(b *testing.B) {
				actual := lookAndSaySequence(tc.sequence)
				if actual != tc.expected {
					b.Fatalf("expected %s, got: %s", tc.expected, actual)
				}
			})
		}
	}
}
