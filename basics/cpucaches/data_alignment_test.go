package cpucaches

import "testing"

// BenchmarkSumFooV1 benchmarks the summing a slice of FooV1 structs to test data alignment
func BenchmarkSumFooV1(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	count := 10_000

	tests := []FooV1{}
	for x := 0; x < count; x++ {
		tests = append(tests, FooV1{i: int64(x), b1: byte(x), b2: byte(x)})
	}

	for i := 0; i < b.N; i++ {
		sumFooV1(tests)
	}
}

// BenchmarkSumFooV2 benchmarks the summing a slice of FooV2 structs to test data alignment in descending order
func BenchmarkSumFooV2(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	count := 10_000

	tests := []FooV2{}
	for x := 0; x < count; x++ {
		tests = append(tests, FooV2{i: int64(x), b1: byte(x), b2: byte(x)})
	}

	for i := 0; i < b.N; i++ {
		sumFooV2(tests)
	}
}
