package rot13

import "testing"

type testCases struct {
	desc     string
	input    string
	expected string
}

func TestRot13(t *testing.T) {
	tests := []testCases{
		{
			desc:     "Should reverse from rot13 to normal text",
			input:    "EBG13 rknzcyr.",
			expected: "ROT13 example.",
		},
		{
			desc:     "Should perform rot13 cypher on normal text",
			input:    "How can you tell an extrovert from an\nintrovert at NSA? Va gur ryringbef,\ngur rkgebireg ybbxf ng gur BGURE thl'f fubrf.",
			expected: "Ubj pna lbh gryy na rkgebireg sebz na\nvagebireg ng AFN? In the elevators,\nthe extrovert looks at the OTHER guy's shoes.",
		},
		{
			desc:     "Should not perform rot13 cypher on numbers",
			input:    "123",
			expected: "123",
		},
		{
			desc:     "Should perform reverse rot13 on text",
			input:    "Guvf vf npghnyyl gur svefg xngn V rire znqr. Gunaxf sbe svavfuvat vg! :)",
			expected: "This is actually the first kata I ever made. Thanks for finishing it! :)",
		},
		{
			desc:     "Should not perform rot13 on text with special characters",
			input:    "@[`{",
			expected: "@[`{",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			actual := Rot13(tt.input)
			if actual != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, actual)
			}
		})
	}
}
