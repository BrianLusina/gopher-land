package natowords

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNatoWords(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NatoWords Suite")
}

var _ = Describe("NatoWords", func() {
	It("should return India Foxtrot , Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta ? for input If, you can read?", func() {
		input := "If, you can read?"
		expected := "India Foxtrot ,Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta ?"
		Expect(toNato(input)).To(Equal(expected))
	})
})
