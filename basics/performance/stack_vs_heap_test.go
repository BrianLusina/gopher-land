package performance

import "testing"

var globalValue int
var globalPtr *int

func BenchmarkSumValue(b *testing.B) {
	b.ReportAllocs() // reports heap allocations
	var local int
	for i := 0; i < b.N; i++ {
		local = sumValue(i, i)
	}
	globalValue = local
}

func BenchmarkSumPtr(b *testing.B) {
	b.ReportAllocs()
	var local *int
	for i := 0; i < b.N; i++ {
		local = sumPtr(i, i)
	}
	globalPtr = local
}
