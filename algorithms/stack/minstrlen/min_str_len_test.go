package minstrlen

import "testing"

type testCase struct {
	input    string
	expected int
}

func getTestCases() []testCase {
	return []testCase{
		{input: "ABCD", expected: 0},
		{input: "ACDBD", expected: 1},
		{input: "ACBD", expected: 4},
		{input: "ABCDXYZCDAB", expected: 3},
		{input: "A", expected: 1},
		{input: "ABFCACDB", expected: 2},
		{input: "ACBBD", expected: 5},
	}
}

func TestMinStrLenStack(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		result := MinStrLenStack(tc.input)
		if result != tc.expected {
			t.Errorf("For input '%s', expected %d but got %d", tc.input, tc.expected, result)
		}
	}
}

func BenchmarkMinStrLenStack(b *testing.B) {
	testCases := getTestCases()
	for b.Loop() {
		for _, tc := range testCases {
			MinStrLenStack(tc.input)
		}
	}
}

func TestMinStrLenStrReplace(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		result := MinStrLenReplace(tc.input)
		if result != tc.expected {
			t.Errorf("For input '%s', expected %d but got %d", tc.input, tc.expected, result)
		}
	}
}

func BenchmarkMinStrLenStrReplace(b *testing.B) {
	testCases := getTestCases()
	for b.Loop() {
		for _, tc := range testCases {
			MinStrLenReplace(tc.input)
		}
	}
}

func TestMinStrLenStrInPlaceManipulation(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		result := MinStrLenInPlaceManipulation(tc.input)
		if result != tc.expected {
			t.Errorf("For input '%s', expected %d but got %d", tc.input, tc.expected, result)
		}
	}
}

func BenchmarkMinStrLenStrInPlaceManipulation(b *testing.B) {
	testCases := getTestCases()
	for b.Loop() {
		for _, tc := range testCases {
			MinStrLenInPlaceManipulation(tc.input)
		}
	}
}
