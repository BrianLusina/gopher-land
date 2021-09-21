package maximumtime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumTime(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(MaximumTime("2?:?0"), "23:50")

	assert.Equal(MaximumTime("0?:3?"), "09:39")

	assert.Equal(MaximumTime("1?:22"), "19:22")

	assert.Equal(MaximumTime("??:3?"), "23:39")
}
