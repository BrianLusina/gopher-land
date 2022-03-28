package railfence

import "sort"

type (
	pair struct {
		index, value int
		letter       byte
	}

	pairs []pair
)

func fencePattern(rails, size int) pairs {
	fence := func(rail, size int) pairs {
		var p pairs
		v, direction := 0, 1
		for i := 0; i < size; i++ {
			p = append(p, pair{index: i, value: v})
			v += direction
			if v == 0 {
				direction = 1
			} else if v == rail-1 {
				direction = -1
			}
		}
		return p
	}

	f := fence(rails, size)
	sort.Slice(f, func(i, j int) bool {
		f1, f2 := f[i], f[j]
		return f1.value < f2.value || (f1.value == f2.value && f1.index < f2.index)
	})
	return f
}

func Encode(message string, rails int) string {
	fence := fencePattern(rails, len(message))
	encoded := []byte{}
	for _, p := range fence {
		encoded = append(encoded, message[p.index])
	}
	return string(encoded)
}

func Decode(message string, rails int) string {
	fence := fencePattern(rails, len(message))

	for i := range message {
		fence[i].letter = message[i]
	}

	sort.Slice(fence, func(i, j int) bool { return fence[i].index < fence[j].index })
	decoded := []byte{}
	for _, p := range fence {
		decoded = append(decoded, p.letter)
	}

	return string(decoded)
}
