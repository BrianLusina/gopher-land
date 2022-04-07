package validbraces

import (
	"fmt"
	"math/rand"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func TestValidBraces(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ValidBraces Test Suite")
}

var _ = Describe("ValidBraces", func() {
	It("should return true for complex latex expression", func() {
		input := "\\left(\\begin{array}{cc} \\frac{1}{3} & x\\\\ \\mathrm{e}^{x} &... x^2 \\end{array}\\right)"
		expected := true
		actual := ValidBraces(input)
		assert.Equal(GinkgoT(), expected, actual)
	})

	It("should return true for valid math expression", func() {
		input := "(((185 + 223.85) * 15) - 543)/2"
		expected := true
		actual := ValidBraces(input)
		assert.Equal(GinkgoT(), expected, actual)
	})

	It("should return true for empty strings", func() {
		input := ""
		expected := true
		actual := ValidBraces(input)
		assert.Equal(GinkgoT(), expected, actual)
	})

	It("should return true for (){}[]", func() {
		input := "(){}[]"
		expected := true
		actual := ValidBraces(input)
		assert.Equal(GinkgoT(), expected, actual)
	})

	It("should return true for ([{}])", func() {
		input := "([{}])"
		expected := true
		actual := ValidBraces(input)
		assert.Equal(GinkgoT(), expected, actual)
	})

	It("should return false for (}", func() {
		input := "(}"
		expected := false
		actual := ValidBraces(input)
		assert.Equal(GinkgoT(), expected, actual)
	})

	It("should return false for [(])", func() {
		input := "[(])"
		expected := false
		actual := ValidBraces(input)
		assert.Equal(GinkgoT(), expected, actual)
	})

	It("should return false for [({)](]", func() {
		input := "[({)](]"
		expected := false
		actual := ValidBraces(input)
		assert.Equal(GinkgoT(), expected, actual)
	})

	It("should return true for {}()[]", func() {
		input := "{}()[]"
		expected := true
		actual := ValidBraces(input)
		Expect(actual).To(Equal(expected))
	})

	It("should return true for ([{}])", func() {
		input := "([{}])"
		expected := true
		actual := ValidBraces(input)
		Expect(actual).To(Equal(expected))
	})

	It("should return false for ([{)]}", func() {
		input := "([{)]}"
		expected := false
		actual := ValidBraces(input)
		Expect(actual).To(Equal(expected))
	})

	It("should return false for ([}{])", func() {
		input := "([}{])"
		expected := false
		actual := ValidBraces(input)
		Expect(actual).To(Equal(expected))
	})

	It("should return true for {}({})[]", func() {
		input := "{}({})[]"
		expected := true
		actual := ValidBraces(input)
		Expect(actual).To(Equal(expected))
	})

	It("should return true for (({{[[]]}}))", func() {
		input := "(({{[[]]}}))"
		expected := true
		actual := ValidBraces(input)
		Expect(actual).To(Equal(expected))
	})

	It("should return false for (((({{}})[]))", func() {
		input := "(((({{}})[]))"
		expected := false
		actual := ValidBraces(input)
		Expect(actual).To(Equal(expected))
	})

	It("should return true for )(}{[]", func() {
		input := ")(}{[]"
		expected := false
		actual := ValidBraces(input)
		Expect(actual).To(Equal(expected))
	})

	It("should return true for (({[]{()}}[{(())}[]]))", func() {
		input := "(({[]{()}}[{(())}[]]))"
		expected := true
		actual := ValidBraces(input)
		Expect(actual).To(Equal(expected))
	})

	parens := []string{"()", "[]", "{}"}

	rng := rand.New(rand.NewSource(rand.Int63()))

	for i := 0; i < 100; i++ {
		str := parens[rng.Int31n(3)]

		for i := rng.Int31n(10); i > 0; i-- {
			index := rng.Int31n(int32(len(str)))
			str = str[:index] + parens[rng.Int31n(3)] + str[index:]
		}

		valid := true
		//flip a coin - valid or invalid?
		if rng.Int31n(2) == 1 {
			//tails, time to invalidate this bitch!
			valid = false
			index := rng.Int31n(int32(len(str)))
			str = str[:index] + str[index+1:]
		}

		It(fmt.Sprintf("Should return %v for %s", valid, str), func() {
			actual := ValidBraces(str)
			assert.Equal(GinkgoT(), valid, actual)
		})
	}
})
