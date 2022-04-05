package worddictionary

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"testing"
)

func TestWordDictionary(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "WordDictionary Suite")
}

var _ = ginkgo.Describe("WordDictionary", func() {
	ginkgo.It("should return correct results for search after adding bad, dad, mad, pad", func() {
		wordDictionary := NewWordDictionary()

		wordDictionary.AddWord("bad")
		wordDictionary.AddWord("dad")
		wordDictionary.AddWord("mad")

		gomega.Expect(wordDictionary.Search("pad")).To(gomega.BeFalse())
		gomega.Expect(wordDictionary.Search("bad")).To(gomega.BeTrue())
		gomega.Expect(wordDictionary.Search("b..")).To(gomega.BeTrue())
		gomega.Expect(wordDictionary.Search(".ad")).To(gomega.BeTrue())
	})
})
