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

func TestEqualUnorderedSlices(t *testing.T) {
	testCases := []struct {
		a, b     []int
		expected bool
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}, true},
		{[]int{1, 2, 2}, []int{2, 1, 1}, false},
		{[]int{1, 2, 3}, []int{4, 5, 6}, false},
		{[]int{1, 2, 3}, []int{1, 2, 3, 4}, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("EqualUnorderedSlices(%v, %v)", tc.a, tc.b), func(t *testing.T) {
			actual := EqualUnorderedSlices(tc.a, tc.b)
			if tc.expected != actual {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
