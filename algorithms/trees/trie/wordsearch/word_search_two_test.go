package wordsearch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type wordSearchTwoTestCase struct {
	grid     [][]string
	words    []string
	expected []string
}

var wordSearchTwoTestCases = []wordSearchTwoTestCase{
	{
		grid: [][]string{
			{"C", "S", "L", "I", "M"},
			{"O", "I", "L", "M", "O"},
			{"O", "L", "I", "E", "O"},
			{"R", "T", "A", "S", "N"},
			{"S", "I", "T", "A", "C"},
		},
		words:    []string{"SLIME", "SAILOR", "MATCH", "COCOON"},
		expected: []string{"SLIME", "SAILOR"},
	},
	{
		grid: [][]string{
			{"C", "S", "L", "I", "M"},
			{"O", "I", "B", "M", "O"},
			{"O", "L", "U", "E", "O"},
			{"N", "L", "Y", "S", "N"},
			{"S", "I", "N", "E", "C"},
		},
		words:    []string{"BUY", "STUFF", "ONLINE", "NOW"},
		expected: []string{"BUY", "ONLINE"},
	},
	{
		grid: [][]string{
			{"C", "O", "L", "I", "M"},
			{"I", "N", "L", "M", "O"},
			{"A", "L", "I", "E", "O"},
			{"R", "T", "A", "S", "N"},
			{"S", "I", "T", "A", "C"},
		},
		words:    []string{"REINDEER", "IN", "RAIN"},
		expected: []string{"IN", "RAIN"},
	},
	{
		grid: [][]string{
			{"P", "S", "L", "A", "M"},
			{"O", "P", "U", "R", "O"},
			{"O", "L", "I", "E", "O"},
			{"R", "T", "A", "S", "N"},
			{"S", "I", "T", "A", "C"},
		},
		words:    []string{"TOURISM", "DESTINED", "POPULAR"},
		expected: []string{"POPULAR"},
	},
	{
		grid: [][]string{
			{"O", "A", "A", "N"},
			{"E", "T", "A", "E"},
			{"I", "H", "K", "R"},
			{"I", "F", "L", "V"},
		},
		words:    []string{"OATH", "PEA", "EAT", "RAIN"},
		expected: []string{"OATH", "EAT"},
	},
}

func TestFindStrings(t *testing.T) {
	for _, tc := range wordSearchTwoTestCases {
		t.Run(fmt.Sprintf("grid: %v, words: %v", tc.grid, tc.words), func(t *testing.T) {
			actual := FindStrings(tc.grid, tc.words)
			if !assert.ElementsMatch(t, tc.expected, actual) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func BenchmarkFindStrings(b *testing.B) {
	for b.Loop() {
		for _, tc := range wordSearchTwoTestCases {
			FindStrings(tc.grid, tc.words)
		}
	}
}
