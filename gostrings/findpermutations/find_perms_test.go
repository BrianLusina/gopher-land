package findpermutations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	input    string
	expected []string
}

var testCases = []testCase{
	{
		name:     "empty string",
		input:    "",
		expected: []string{},
	},
	{
		name:     "ABC",
		input:    "ABC",
		expected: []string{"ABC", "ACB", "BAC", "BCA", "CAB", "CBA"},
	},
	{
		name:  "ABSG",
		input: "ABSG",
		expected: []string{
			"ABGS", "ABSG", "AGBS", "AGSB", "ASBG", "ASGB", "BAGS",
			"BASG", "BGAS", "BGSA", "BSAG", "BSGA", "GABS", "GASB",
			"GBAS", "GBSA", "GSAB", "GSBA", "SABG", "SAGB", "SBAG",
			"SBGA", "SGAB", "SGBA",
		},
	},
}

func TestFindPerms(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := FindPermutations(tc.input)
			if !assert.Equal(t, tc.expected, actual) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func BenchmarkFindPerms(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark")
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindPermutations(tc.input)
			}
		})
	}
}
