package gasstations

import "testing"

type testCase struct {
	gas      []int
	cost     []int
	expected int
}

var testCases = []testCase{
	{
		gas:      []int{1, 2, 3, 4, 5},
		cost:     []int{3, 4, 5, 1, 2},
		expected: 3,
	},
	{
		gas:      []int{2, 3, 4},
		cost:     []int{3, 4, 3},
		expected: -1,
	},
}

func TestCanCompleteCircuit(t *testing.T) {
	for _, tc := range testCases {
		actual := canCompleteCircuit(tc.gas, tc.cost)
		if actual != tc.expected {
			t.Fatalf("canCompleteCircuit(gas=%v, cost=%v)=%d, expected = %d", tc.gas, tc.cost, actual, tc.expected)
		}
	}
}

func BenchmarkCanCompleteCircuit(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := canCompleteCircuit(tc.gas, tc.cost)
			if actual != tc.expected {
				b.Fatalf("canCompleteCircuit(gas=%v, cost=%v)=%d, expected = %d", tc.gas, tc.cost, actual, tc.expected)
			}
		}
	}
}
