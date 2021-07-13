package isomorphic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsomorphicStrings(t *testing.T) {
	assert := assert.New(t)

	assert.False(IsIsomorphic("aaab", "b"))

	assert.True(IsIsomorphic("aab", "xxy"))

	assert.False(IsIsomorphic("aab", "xyz"))

	assert.True(IsIsomorphic("egg", "add"))

	assert.False(IsIsomorphic("badc", "baba"))
}
