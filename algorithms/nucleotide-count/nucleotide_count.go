package dna

import (
	"errors"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[byte]int

// DNA is a list of nucleotides
type DNA string

const nucleotides = "ACGT"

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (dna DNA) Counts() (Histogram, error) {
	var totalCount int

	histogram := make(Histogram)

	for index := range nucleotides {
		nucleotide := nucleotides[index]
		histogram[nucleotide] = strings.Count(string(dna), string(nucleotide))
		totalCount += histogram[nucleotide]
	}

	if totalCount != len(dna) {
		return nil, errors.New("dna strand contains invalid nucleotide")
	}

	return histogram, nil
}
