package happynumber

import (
	"fmt"
	"testing"
)

type testCase struct {
	n        int
	expected bool
}

var testCases = []testCase{
	{n: 19, expected: true},
	{n: 2, expected: false},
	{n: 7, expected: true},
	{n: 20, expected: false},
	{n: 23, expected: true},
	{n: 2147483646, expected: false},
	{n: 1, expected: true},
	{n: 8, expected: false},
	{n: 5, expected: false},
	{n: 25, expected: false},
}

func TestIsHappyNumber(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("IsHappyNumber(%d)", tc.n), func(t *testing.T) {
			actual := IsHappyNumber(tc.n)
			if actual != tc.expected {
				t.Fail()
			}
		})
	}
}

func BenchmarkIsHappyNumber(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}

	for b.Loop() {
		for _, tc := range testCases {
			IsHappyNumber(tc.n)
		}
	}
}

func TestIsHappyNumberFastAndSlow(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("IsHappyNumber(%d)", tc.n), func(t *testing.T) {
			actual := IsHappyNumberFastAndSlow(tc.n)
			if actual != tc.expected {
				t.Fail()
			}
		})
	}
}

func BenchmarkIsHappyNumberFastAndSlow(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}

	for b.Loop() {
		for _, tc := range testCases {
			IsHappyNumberFastAndSlow(tc.n)
		}
	}
}
