package islongpressed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	description string
	name        string
	typed       string
	expected    bool
}{
	{
		description: "Should return true for name alex and typed aaleex",
		name:        "alex",
		typed:       "aaleex",
		expected:    true,
	},
	{
		description: "Should return true for name saeed and typed ssaaedd",
		name:        "saeed",
		typed:       "ssaaedd",
		expected:    false,
	},
}

func TestIsLongPressed(t *testing.T) {
	for _, test := range testCases {
		actual := isLongPressedName(test.name, test.typed)
		assert.Equal(t, test.expected, actual)
	}
}
