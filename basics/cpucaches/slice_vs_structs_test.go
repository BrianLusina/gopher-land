package cpucaches

import "testing"

func BenchmarkSumFoo(b *testing.B) {
	foos := []Foo{}

	for i := 1; i < 16; i++ {
		f := Foo{
			a: i,
			b: i,
		}
		foos = append(foos, f)
	}

	for i := 0; i < b.N; i++ {
		sumFoo(foos)
	}
}

// This is faster than a slice of structs
func BenchmarkSumBar(b *testing.B) {
	bar := Bar{
		a: []int{},
		b: []int{},
	}

	for i := 1; i < 16; i++ {
		bar.a = append(bar.a, i)
		bar.b = append(bar.b, i)
	}

	for i := 0; i < b.N; i++ {
		sumBar(bar)
	}
}
