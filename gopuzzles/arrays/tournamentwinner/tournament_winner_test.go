package tournamentwinner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name         string
	competitions [][]string
	results      []int
	expected     string
}

var testCases = []testCase{
	{
		name: "Test case 1",
		competitions: [][]string{
			{"HTML", "C#"},
			{"C#", "Python"},
			{"Python", "HTML"},
		},
		results: []int{
			0,
			0,
			1,
		},
		expected: "Python",
	},
	{
		name: "Test case 2",
		competitions: [][]string{
			{"HTML", "Java"},
			{"Java", "Python"},
			{"Python", "HTML"},
		},
		results: []int{
			0,
			1,
			1,
		},
		expected: "Java",
	},
}

func TestTournamentWinner(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tournamentWinner(tc.competitions, tc.results)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkTournamentWinner(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for b.Loop() {
		for _, tc := range testCases {
			tournamentWinner(tc.competitions, tc.results)
		}
	}
}

func TestTournamentWinnerV2(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tournamentWinnerV2(tc.competitions, tc.results)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkTournamentWinnerV2(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for b.Loop() {
		for _, tc := range testCases {
			tournamentWinnerV2(tc.competitions, tc.results)
		}
	}
}
