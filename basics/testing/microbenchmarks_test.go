package gotesting

import (
	"sync/atomic"
	"testing"
)

/**
* These tests can be run with the command go test -bench=. -count=10 | tee stats.txt
This will generate the stats.txt file and used to evaluate and eliminate the assumption that one function is faster than the other by using the benchstat tool
which can be found here https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
Running that tool like below benchstat stats.txt
**/

func BenchmarkAtomicStoreInt32(b *testing.B) {
	var v int32
	for i := 0; i < b.N; i++ {
		atomic.StoreInt32(&v, 1)
	}
}

func BenchmarkAtomicStoreInt64(b *testing.B) {
	var v int64
	for i := 0; i < b.N; i++ {
		atomic.StoreInt64(&v, 1)
	}
}
