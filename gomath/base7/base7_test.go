package base7

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestBase7(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PrimeGap Suite")
}

var _ = Describe("Base7", func() {
	It("should return 202 with input of 100", func() {
		input := 100
		expected := "202"
		actual := ConvertToBase7(input)
		Expect(actual).To(Equal(expected))
	})
})
