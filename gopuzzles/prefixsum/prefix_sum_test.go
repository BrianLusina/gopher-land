package prefixsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	description string
	input       []int
	expected    []int
	expectedSum int
}

var testCases = []testCase{
	{
		description: "prefix sum of empty array",
		input:       []int{},
		expected:    []int{},
		expectedSum: 0,
	},
	{
		description: "prefix sum of single element array",
		input:       []int{1},
		expected:    []int{1},
		expectedSum: 1,
	},
	{
		description: "prefix sum of two element array",
		input:       []int{1, 2},
		expected:    []int{1, 3},
		expectedSum: 3,
	},
	{
		description: "prefix sum of three element array",
		input:       []int{1, 2, 3},
		expected:    []int{1, 3, 6},
		expectedSum: 6,
	},
	{
		description: "prefix sum of four element array",
		input:       []int{1, 2, 3, 4},
		expected:    []int{1, 3, 6, 10},
		expectedSum: 10,
	},
}

func TestPrefixSum(t *testing.T) {
	for _, testCase := range testCases {
		output, sum := PrefixSum(testCase.input)
		if sum != testCase.expectedSum {
			t.Fatalf("expected sum %d, got %d", testCase.expectedSum, sum)
		}
		if !assert.EqualValues(t, testCase.expected, output) {
			t.Errorf("%s: expected %v, got %v", testCase.description, testCase.expected, output)
		}
	}
}

func BenchmarkPrefixSum(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark")
	}
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			PrefixSum(testCase.input)
		}
	}
}
