package gotesting

import "testing"

func BenchmarkPopcnt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcont(uint64(i))
	}
}

var global uint64

func BenchmarkPopcnt2(b *testing.B) {
	var v uint64

	for i := 0; i < b.N; i++ {
		v = popcont(uint64(i))
	}
	global = v
}
