package asteroidcollision

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	asteroids []int
	expected  []int
}

var testCases = []testCase{
	{
		asteroids: []int{5, 10, -5},
		expected:  []int{5, 10},
	},
	{
		asteroids: []int{8, -8},
		expected:  []int{},
	},
	{
		asteroids: []int{10, 2, -5},
		expected:  []int{10},
	},
}

func TestAsteroidCollision(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("asteroidCollision(%v)", tc.asteroids), func(t *testing.T) {
			actual := asteroidCollision(tc.asteroids)
			assert.ElementsMatch(t, actual, tc.expected)
		})
	}
}

func BenchmarkAsteroidCollision(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("asteroidCollision(%v)", tc.asteroids), func(b *testing.B) {
				actual := asteroidCollision(tc.asteroids)
				assert.ElementsMatch(b, actual, tc.expected)
			})
		}
	}
}
