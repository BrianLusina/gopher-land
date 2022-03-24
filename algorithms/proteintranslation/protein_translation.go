package protein

import "errors"

var codonToRnaMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// ErrStop thrown when STOP codon is encountered
var ErrStop = errors.New("STOP codon encountered")

// ErrInvalidBase thrown when there is an invalid codon
var ErrInvalidBase = errors.New("invalid codon encountered")

// FromCodon translates a codon to rna
func FromCodon(codon string) (rna string, err error) {
	rna, ok := codonToRnaMap[codon]
	if !ok {
		rna = ""
		err = ErrInvalidBase
	}
	if rna == "STOP" {
		err = ErrStop
		rna = ""
	}
	return
}

// FromRNA translates an RNA to codon
func FromRNA(rna string) (proteins []string, err error) {

	for i := 0; i < len(rna); i += 3 {
		codon := rna[i : i+3]
		protein, err := FromCodon(codon)
		switch err {
		case ErrStop:
			return proteins, nil
		case ErrInvalidBase:
			return proteins, err
		default:
			proteins = append(proteins, protein)
		}
	}
	return
}
