package anagram

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func equal(a, b []string) bool {
	if len(b) != len(a) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

func TestIsAnagram(t *testing.T) {
	for _, tt := range isAnagramTestCases {
		t.Run(fmt.Sprintf("%s and %s", tt.word1, tt.word2), func(t *testing.T) {
			actual := IsAnagram(tt.word1, tt.word2)
			if actual != tt.expected {
				t.Fatalf("expected %v, got: %v", tt.expected, actual)
			}
		})
	}
}

func BenchmarkIsAnagram(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range isAnagramTestCases {
			IsAnagram(tt.word1, tt.word2)
		}
	}
}

func TestDetectAnagrams(t *testing.T) {
	for _, tt := range testCases {
		actual := Detect(tt.subject, tt.candidates)
		if !equal(tt.expected, actual) {
			msg := `FAIL: %s
	Subject %s
	Candidates %q
	Expected %q
	Got %q
				`
			t.Fatalf(msg, tt.description, tt.subject, tt.candidates, tt.expected, actual)
		} else {
			t.Logf("PASS: %s", tt.description)
		}
	}
}

func BenchmarkDetectAnagrams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			Detect(tt.subject, tt.candidates)
		}
	}
}

// TODO: sort a slice of slice for tests to pass
// func TestGroupAnagrams(t *testing.T) {
// 	for _, tc := range groupAnagramTestCases {

// 		actual := GroupAnagrams(tc.strs)

// 		if !reflect.DeepEqual(actual, tc.expected) {
// 			t.Fatalf("GroupAnagrams(%v) = %v, expected: %v", tc.strs, actual, tc.expected)
// 		} else {
// 			t.Logf("Pass: Actual = %v, Expected = %v", actual, tc.expected)
// 		}
// 	}
// }

// func BenchmarkGroupAnagrams(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		for _, tt := range groupAnagramTestCases {
// 			GroupAnagrams(tt.strs)
// 		}
// 	}
// }

func TestGroupAnagrams(t *testing.T) {
	for _, tc := range generateAnagramTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := GenerateAnagrams(tc.word)
			if !assert.Equal(t, tc.expected, actual) {
				msg := `FAIL: %s
				Word %s
				Expected %q
				Got %q
							`
				t.Fatalf(msg, tc.word, tc.expected, actual)
			} else {
				t.Logf("PASS: %s", tc.word)
			}

		})
	}
}

func BenchmarkGroupAnagrams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range generateAnagramTestCases {
			GenerateAnagrams(tc.word)
		}
	}
}
