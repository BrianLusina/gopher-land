package strand

var rnaTests = []struct {
	description string
	input       string
	expected    string
}{
	{"Empty RNA sequence",
		"", "",
	},
	{"RNA complement of cytosine is guanine",
		"C", "G",
	},
	{"RNA complement of guanine is cytosine",
		"G", "C",
	},
	{"RNA complement of thymine is adenine",
		"T", "A",
	},
	{"RNA complement of adenine is uracil",
		"A", "U",
	},
	{"RNA complement",
		"ACGTGGTCTTAA", "UGCACCAGAAUU",
	},
}
