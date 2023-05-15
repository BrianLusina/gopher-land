package climbstairs

import "testing"

type testCase struct {
	name     string
	stairs   int
	expected int
}

var testCases = []testCase{
	{
		name:     "Number of ways of going up 1 stair going up either 1 or 2 steps at a time",
		stairs:   1,
		expected: 1,
	},
	{
		name:     "Number of ways of going up 2 stairs going up either 1 or 2 steps at a time",
		stairs:   2,
		expected: 2,
	},
	{
		name:     "Number of ways of going up 3 stairs going up either 1 or 2 steps at a time",
		stairs:   3,
		expected: 3,
	},
}

func TestClimbStairs(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := ClimbStairs(tc.stairs)
			if actual != tc.expected {
				t.Fatalf("ClimbStairs(%d) = %d, expected=%d", tc.stairs, actual, tc.expected)
			}
		})
	}
}

func BenchmarkClimbStairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(tc.name, func(b *testing.B) {
				actual := ClimbStairs(tc.stairs)
				if actual != tc.expected {
					b.Fatalf("ClimbStairs(%d) = %d, expected=%d", tc.stairs, actual, tc.expected)
				}
			})
		}
	}
}

func TestClimbStaircase(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := ClimbStaircase(tc.stairs)
			if actual != tc.expected {
				t.Fatalf("ClimbStairs(%d) = %d, expected=%d", tc.stairs, actual, tc.expected)
			}
		})
	}
}

func BenchmarkClimbStaircase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(tc.name, func(b *testing.B) {
				actual := ClimbStaircase(tc.stairs)
				if actual != tc.expected {
					b.Fatalf("ClimbStairs(%d) = %d, expected=%d", tc.stairs, actual, tc.expected)
				}
			})
		}
	}
}
