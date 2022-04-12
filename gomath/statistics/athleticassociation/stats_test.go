package athleticassociation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestStats(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stats Suite")
}

var _ = Describe("Stati", func() {

	It("should handle input of '' & return ''", func() {
		input := ""
		actual := Stati(input)
		expected := ""
		Expect(actual).To(Equal(expected))
	})

	It("should handle input of '01|15|59, 1|47|16, 01|17|20, 1|32|34, 2|17|17' & return 'Range: 01|01|18 Average: 01|38|05 Median: 01|32|34'", func() {
		input := "01|15|59, 1|47|16, 01|17|20, 1|32|34, 2|17|17"
		actual := Stati(input)
		expected := "Range: 01|01|18 Average: 01|38|05 Median: 01|32|34"
		Expect(actual).To(Equal(expected))
	})

	It("should handle input of '02|15|59, 2|47|16, 02|17|20, 2|32|34, 2|17|17, 2|22|00, 2|31|41' and return 'Range: 00|31|17 Average: 02|26|18 Median: 02|22|00'", func() {
		input := "02|15|59, 2|47|16, 02|17|20, 2|32|34, 2|17|17, 2|22|00, 2|31|41"
		actual := Stati(input)
		expected := "Range: 00|31|17 Average: 02|26|18 Median: 02|22|00"
		Expect(actual).To(Equal(expected))
	})
})
