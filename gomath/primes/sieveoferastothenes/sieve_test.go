package sieveoferastothenes

import (
	"reflect"
	"testing"
)

func TestSieve(t *testing.T) {
	for _, tc := range testCases {
		p := Sieve(2, tc.limit)
		if len(p) != 0 || len(tc.expected) != 0 {
			if !reflect.DeepEqual(p, tc.expected) {
				t.Fatalf("FAIL: %s\nSieve(%d)\nExpected %v\nActual  %v",
					tc.description, tc.limit, tc.expected, p)
			}
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkSieve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Sieve(2, tc.limit)
		}
	}
}
