package cpucaches

import "testing"

func BenchmarkFalseSharingWithoutPadding(b *testing.B) {
	inputs := []Input{}
	limit := 100
	for i := 0; i < limit; i++ {
		input := Input{
			a: int64(i),
			b: int64(i),
		}
		inputs = append(inputs, input)
	}

	for x := 0; x < b.N; x++ {
		countWithoutPadding(inputs)
	}
}

func BenchmarkFalseSharingWithPadding(b *testing.B) {
	inputs := []Input{}
	limit := 100
	for i := 0; i < limit; i++ {
		input := Input{
			a: int64(i),
			b: int64(i),
		}
		inputs = append(inputs, input)
	}

	for x := 0; x < b.N; x++ {
		countWithPadding(inputs)
	}
}
