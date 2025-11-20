package utils

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {
	type testCase struct {
		n        int
		expected int
	}

	absTestCases := []testCase{
		{
			n:        -1,
			expected: 1,
		},
		{
			n:        1,
			expected: 1,
		},
		{
			n:        0,
			expected: 0,
		},
	}

	for _, tc := range absTestCases {
		t.Run(fmt.Sprintf("Abs(%d)", tc.n), func(t *testing.T) {
			actual := Abs(tc.n)
			if tc.expected != actual {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestSumOfSquaredDigits(t *testing.T) {
	type testCase struct {
		n        int
		expected int
	}

	testCases := []testCase{{
		n:        19,
		expected: 82,
	}}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("SumOfSquaredDigits(%d)", tc.n), func(t *testing.T) {
			actual := SumOfSquaredDigits(tc.n)
			if actual != tc.expected {
				t.Fail()
			}
		})
	}
}
