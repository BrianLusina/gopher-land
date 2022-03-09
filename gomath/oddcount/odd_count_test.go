package oddcount

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

var oddCountTests = []struct {
	description string
	n           int
	expected    int
}{
	{
		description: "limit of 15",
		n:           15,
		expected:    7,
	},
	{
		description: "limit of 15023",
		n:           15023,
		expected:    7511,
	},
}

func TestOddCount(t *testing.T) {
	for _, test := range oddCountTests {
		if actual := OddCount(test.n); actual != test.expected {
			t.Fatalf("FAIL: OddCount(%d): expected %d, actual %d", test.n, test.expected, actual)
		} else {
			t.Logf("PASS: OddCount(%d) %d", test.n, actual)
		}
	}
}

func TestOddCountRandom(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i <= 100; i++ {
		n := random(0, math.MaxInt64)
		actual := (OddCount(n))
		expected := n / 2
		if actual != expected {
			t.Fatalf("FAIL: OddCount(%d): expected %d, actual %d", n, expected, actual)
		} else {
			t.Logf("PASS: OddCount(%d) %d", n, actual)
		}
	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
