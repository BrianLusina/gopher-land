package removingstars

import (
	"fmt"
	"testing"
)

type testCase struct {
	word, expected string
}

var testCases = []testCase{
	{
		word:     "leet**cod*e",
		expected: "lecoe",
	},
	{
		word:     "erase*****",
		expected: "",
	},
}

func TestRemoveStars(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("removeStars(%s)", tc.word), func(t *testing.T) {
			actual := removeStars(tc.word)
			if actual != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, actual)
			}
		})
	}
}

func TestRemoveStars2(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("removeStars2(%s)", tc.word), func(t *testing.T) {
			actual := removeStars2(tc.word)
			if actual != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, actual)
			}
		})
	}
}

func BenchmarkRemoveStars(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("removeStars(%s)", tc.word), func(b *testing.B) {
				actual := removeStars(tc.word)
				if actual != tc.expected {
					b.Errorf("expected %s, got %s", tc.expected, actual)
				}
			})
		}
	}
}

func BenchmarkRemoveStars2(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("removeStars2(%s)", tc.word), func(b *testing.B) {
				actual := removeStars2(tc.word)
				if actual != tc.expected {
					b.Errorf("expected %s, got %s", tc.expected, actual)
				}
			})
		}
	}
}
