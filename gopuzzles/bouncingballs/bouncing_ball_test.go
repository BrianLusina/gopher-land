package bouncingballs

import "testing"

type testCase struct {
	h        float64
	bounce   float64
	window   float64
	expected int
}

var testCases = []testCase{
	{
		h:        3,
		bounce:   0.66,
		window:   1.5,
		expected: 3,
	},
	{
		h:        40,
		bounce:   0.4,
		window:   10,
		expected: 3,
	},
	{
		h:        10,
		bounce:   0.6,
		window:   10,
		expected: -1,
	},
	{
		h:        40,
		bounce:   1,
		window:   10,
		expected: -1,
	},
	{
		h:        5,
		bounce:   -1,
		window:   1.5,
		expected: -1,
	},
}

func TestBouncingBall(t *testing.T) {
	for _, tc := range testCases {
		actual := bouncingBall(tc.h, tc.bounce, tc.window)

		if actual != tc.expected {
			t.Fatalf("bouncingBall(%f, %f, %f) = %d expected %d", tc.h, tc.bounce, tc.window, actual, tc.expected)
		}
	}
}

func BenchmarkBouncingBall(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			bouncingBall(tc.h, tc.bounce, tc.window)
		}
	}
}
