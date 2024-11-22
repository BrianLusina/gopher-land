package cpucaches

import "testing"

func BenchmarkAdd(b *testing.B) {
	input := [2]int64{1, 2}
	for x := 0; x < b.N; x++ {
		add(input)
	}
}

func BenchmarkAdd2(b *testing.B) {
	input := [2]int64{1, 2}
	for x := 0; x < b.N; x++ {
		add2(input)
	}
}
