package strand

var dnaRnaMap = map[rune]string{
	'A': "U",
	'T': "A",
	'C': "G",
	'G': "C",
}

// ToRNA converts each nucleotide DNA sequence to its complimentary RNA sequence
func ToRNA(dna string) (rnaStrand string) {

	for _, dnaStrand := range dna {
		rnaStrand += dnaRnaMap[dnaStrand]
	}

	return rnaStrand
}
