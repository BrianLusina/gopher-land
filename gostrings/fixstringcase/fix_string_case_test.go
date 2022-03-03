package fixstringcase

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFixStringCase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FixStringCase Suite")
}

var _ = Describe("FixStringCase", func() {
	It("Should return the correct values for the sample test cases!", func() {
		Expect(fixstringcase("a")).To(Equal("a"))
		Expect(fixstringcase("Z")).To(Equal("Z"))
		Expect(fixstringcase("coDe")).To(Equal("code"))
		Expect(fixstringcase("CODe")).To(Equal("CODE"))
		Expect(fixstringcase("coDE")).To(Equal("code"))
		Expect(fixstringcase("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")).To(Equal("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"))
	})
})
