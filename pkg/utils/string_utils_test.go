package utils

import (
	"fmt"
	"testing"
)

type isAlphanumericTestCase struct {
	char     byte
	expected bool
}

var isAlphanumericTestCases = []isAlphanumericTestCase{
	{
		char:     'a',
		expected: true,
	},
	{
		char:     '/',
		expected: false,
	},
	{
		char:     '0',
		expected: true,
	},
}

func TestIsAlphanumeric(t *testing.T) {
	for _, tc := range isAlphanumericTestCases {
		t.Run(fmt.Sprintf("IsAlphanumeric(%b)", tc.char), func(t *testing.T) {
			actual := IsAlphanumeric(tc.char)
			if tc.expected != actual {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func BenchmarkIsAlphanumeric(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range isAlphanumericTestCases {
			b.Run(fmt.Sprintf("IsAlphanumeric(%b)", tc.char), func(b *testing.B) {
				actual := IsAlphanumeric(tc.char)
				if tc.expected != actual {
					b.Fatalf("expected %v, got %v", tc.expected, actual)
				}
			})
		}
	}
}

type reverseTestCase struct {
	s        string
	expected string
}

var reverseTestCases = []reverseTestCase{
	{
		s:        "hello",
		expected: "olleh",
	},
	{
		s:        "racecar",
		expected: "racecar",
	},
	{
		s:        "world",
		expected: "dlrow",
	},
}

func TestReverse(t *testing.T) {
	for _, tc := range reverseTestCases {
		t.Run(fmt.Sprintf("Reverse(%s)", tc.s), func(t *testing.T) {
			actual := Reverse(tc.s)
			if tc.expected != actual {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range reverseTestCases {
			b.Run(fmt.Sprintf("Reverse(%s)", tc.s), func(b *testing.B) {
				actual := Reverse(tc.s)
				if tc.expected != actual {
					b.Fatalf("expected %v, got %v", tc.expected, actual)
				}
			})
		}
	}
}
