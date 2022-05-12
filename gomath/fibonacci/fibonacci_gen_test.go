package fibonacci

import "testing"

type testCase struct {
	description string
	input       int
	expected    int
}

var testCases = []testCase{
	{
		description: "Get the 10th fibonacci item",
		input:       10,
		expected:    55,
	},
}

func TestNthFibonacci(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := NthFibonacci(testCase.input)
			if actual != testCase.expected {
				t.Fatalf("expected %d, got %d", testCase.expected, actual)
			}
			t.Logf("%dth fibonacci item is %d", testCase.input, actual)
		})
	}
}

func BenchmarkNthFibonacci(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark")
	}
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			NthFibonacci(testCase.input)
		}
	}
}
