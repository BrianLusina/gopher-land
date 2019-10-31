// Calculates the Hamming distance between 2 DNA strands
package hamming

import "errors"

func Distance(a, b string) (count int, err error) {
	if len(a) != len(b) {
		return 0, errors.New("DNA strands of unequal length")
	}

	for x := 0; x < len(a); x++ {
		if a[x] != b[x] {
			count++
		}
	}

	return
}
