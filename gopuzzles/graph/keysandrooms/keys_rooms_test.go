package keysandrooms

import (
	"fmt"
	"testing"
)

type testCase struct {
	rooms    [][]int
	expected bool
}

var testCases = []testCase{
	{
		rooms:    [][]int{{1}, {2}, {3}, {}},
		expected: true,
	},
	{
		rooms:    [][]int{{1, 3}, {3, 0, 1}, {2}, {0}},
		expected: false,
	},
	{
		rooms:    [][]int{{2, 3}, {}, {2}, {1, 3}},
		expected: true,
	},
}

func TestCanVisitAllRooms(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("canVisitAllRooms(%v)", tc.rooms), func(t *testing.T) {
			actual := canVisitAllRooms(tc.rooms)
			if actual != tc.expected {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func BenchmarkCanVisitAllRooms(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("canVisitAllRooms(%v)", tc.rooms), func(b *testing.B) {
				actual := canVisitAllRooms(tc.rooms)
				if actual != tc.expected {
					b.Fatalf("expected %v, got %v", tc.expected, actual)
				}
			})
		}
	}
}
