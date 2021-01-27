package prime

// Nth determines the nth prime in a sequence, starting with 2
func Nth(n int) (nthPrime int, success bool) {
	if n < 0 {
		return 0, false
	}

	switch {
	case n < 1:
		return 0, false
	case n == 1:
		return 2, true
	}

	n--
	nthPrime = 3
	inc := 1
	sqr := 1
	sqrt := 1
	for {
		for f := 3; ; f += 2 {
			if f > sqrt {
				n--
				if n == 0 {
					return nthPrime, true
				}
				break
			}
			if nthPrime%f == 0 {
				break
			}
		}
		nthPrime += 2
		if nthPrime > sqr {
			inc += 2
			sqr += inc
			sqrt++
		}
	}
}
