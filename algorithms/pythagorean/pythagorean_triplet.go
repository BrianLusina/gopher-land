package pythagorean

// Triplet represents the Pythagorean triplets a, b, c
type Triplet [3]int

// isPythagoreanTriplet takes in 3 integers & checks if they are Pythagorean triplets
func isPythagoreanTriplet(a, b, c int) bool {
	return a*a+b*b == c*c
}

// Range returns a list of all Pythagorean triplets with sides in the range min to max inclusive
func Range(min, max int) (ts []Triplet) {
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if isPythagoreanTriplet(a, b, c) {
					ts = append(ts, Triplet{a, b, c})
				}
			}
		}
	}
	return
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c (the perimeter) is equal to p
func Sum(p int) (ts []Triplet) {
	max := p / 2

	for a := 1; a <= max; a++ {
		for b := a; b <= max; b++ {
			if c := p - a - b; isPythagoreanTriplet(a, b, c) {
				ts = append(ts, Triplet{a, b, c})
			}
		}
	}
	return
}
