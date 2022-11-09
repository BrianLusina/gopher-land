package encodedecode

import (
	"testing"

	. "github.com/onsi/gomega"
)

type testCase struct {
	words    []string
	expected []string
}

var testCases = []testCase{
	{
		words:    []string{"lint", "code", "love", "you"},
		expected: []string{"lint", "code", "love", "you"},
	},
	{
		words:    []string{"we", "say", ":", "yes"},
		expected: []string{"we", "say", ":", "yes"},
	},
}

func TestEncodeDecode(t *testing.T) {
	RegisterTestingT(t)
	for _, tc := range testCases {
		actualEncoded := encode(tc.words)
		actualDecoded := decode(actualEncoded)

		Expect(actualDecoded).To(Equal(tc.expected))
	}
}
