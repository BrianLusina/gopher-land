package powerfulintegers

import "testing"

type testCase struct {
	x      int
	y      int
	bound  int
	output []int
}

var testCases = []testCase{
	{2, 3, 10, []int{2, 3, 4, 5, 7, 9, 10}},
	{3, 5, 15, []int{2, 4, 6, 8, 10, 14}},
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[int]bool)
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b {
		if !m[v] {
			return false
		}
	}
	return true
}

func TestPowerfulIntegers(t *testing.T) {
	for _, tc := range testCases {
		result := powerfulIntegers(tc.x, tc.y, tc.bound)
		if !equalSlices(result, tc.output) {
			t.Errorf("For x=%d, y=%d, bound=%d, expected %v but got %v", tc.x, tc.y, tc.bound, tc.output, result)
		}
	}
}

func BenchmarkPowerfulIntegers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		powerfulIntegers(2, 3, 10)
	}
}
