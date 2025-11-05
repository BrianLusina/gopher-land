package repeateddnasequences

import (
	"fmt"
	"gopherland/pkg/utils"
	"testing"
)

type testCase struct {
	dnaSequence    string
	expectedResult []string
}

var testCases = []testCase{
	{
		dnaSequence:    "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
		expectedResult: []string{"AAAAACCCCC", "CCCCCAAAAA"},
	},
	{
		dnaSequence:    "AAAAAAAAAAAAA",
		expectedResult: []string{"AAAAAAAAAA"},
	},
	{
		dnaSequence:    "GGGGGGGGGGGGGGGGGGGG",
		expectedResult: []string{"GGGGGGGGGG"},
	},
	{
		dnaSequence:    "TTTGGGAAATTTGGGAAACC",
		expectedResult: []string{},
	},
	{
		dnaSequence:    "AAAAAAAAAAAAA",
		expectedResult: []string{"AAAAAAAAAA"},
	},
	{
		dnaSequence:    "ACGTACGTACGTACGTACGTACGTACGTACGT",
		expectedResult: []string{"ACGTACGTAC", "CGTACGTACG", "GTACGTACGT", "TACGTACGTA"},
	},
	{
		dnaSequence:    "GTACGTACGTACGCCCCCCCCGGGGG",
		expectedResult: []string{},
	},
}

func TestFindRepeatedDnaSequences(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("dnaSequence=%s", tc.dnaSequence), func(t *testing.T) {
			actualResult := findRepeatedDnaSequences(tc.dnaSequence)
			if !utils.EqualUnorderedSlices(actualResult, tc.expectedResult) {
				t.Errorf("expected %v, but got %v", tc.expectedResult, actualResult)
			}
		})
	}
}

func BenchmarkFindRepeatedDnaSequences(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for b.Loop() {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("dnaSequence=%s", tc.dnaSequence), func(b *testing.B) {
				actualResult := findRepeatedDnaSequences(tc.dnaSequence)
				if !utils.EqualUnorderedSlices(actualResult, tc.expectedResult) {
					b.Errorf("expected %v, but got %v", tc.expectedResult, actualResult)
				}
			})
		}
	}
}
