package prefix

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestTrie(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Trie Suite")
}

var _ = Describe("Trie", func() {
	It("should correctly work for insert(apple) -> search(apple) -> startsWith(app) -> insert(app) -> search(app)", func() {
		trie := NewTrie()
		trie.Insert("apple")

		// return True
		Expect(trie.Search("apple")).To(BeTrue())

		// return False
		Expect(trie.Search("app")).To(BeFalse())

		// return True
		Expect(trie.StartsWith("app")).To(BeTrue())

		trie.Insert("app")

		// return true
		Expect(trie.Search("app")).To(BeTrue())
	})

})
