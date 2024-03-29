package reversewords

import "testing"

type testCase struct {
	name     string
	input    string
	expected string
}

var testCases = []testCase{
	{
		name:     "empty string",
		input:    "",
		expected: "",
	},
	{
		name:     "single word",
		input:    "hello",
		expected: "hello",
	},
	{
		name:     "2 words",
		input:    "hello.world",
		expected: "world.hello",
	},
	{
		name:     "i.like.this.program.very.much",
		input:    "i.like.this.program.very.much",
		expected: "much.very.program.this.like.i",
	},
	{
		name:     "pqr.mno",
		input:    "pqr.mno",
		expected: "mno.pqr",
	},
}

var testCasesWordsInString = []testCase{
	{
		name:     "empty string",
		input:    "",
		expected: "",
	},
	{
		name:     "single word",
		input:    "hello",
		expected: "hello",
	},
	{
		name:     "2 words",
		input:    "hello world",
		expected: "world hello",
	},
	{
		name:     "i like this.program very much",
		input:    "i like this program very much",
		expected: "much very program this like i",
	},
	{
		name:     "pqr mno",
		input:    "pqr mno",
		expected: "mno pqr",
	},
	{
		name:     "the sky is blue",
		input:    "the sky is blue",
		expected: "blue is sky the",
	},
	{
		name:     "  hello world  ",
		input:    "  hello world  ",
		expected: "world hello",
	},
	{
		name:     "a good   example",
		input:    "a good   example",
		expected: "example good a",
	},
}

func TestReverseWords(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := reverseWords(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, actual)
			}
		})
	}
}

func TestReverseWordsInString(t *testing.T) {
	for _, tc := range testCasesWordsInString {
		t.Run(tc.name, func(t *testing.T) {
			actual := reverseWordsInString(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, actual)
			}
		})
	}
}

func BenchmarkReverseWords(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode")
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			reverseWords(tc.input)
		}
	}
}

func BenchmarkReverseWordsInString(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode")
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCasesWordsInString {
			reverseWordsInString(tc.input)
		}
	}
}
