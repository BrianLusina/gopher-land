package spreadsheetencoding

import (
	"fmt"
	"testing"
)

type testCase struct {
	columnId string
	expected int
}

var testCases = []testCase{
	{
		columnId: "ZZ",
		expected: 702,
	},
}

func TestSpreadsheetEncodeColumn(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("spreadsheetEncodeColumn(%s)", tc.columnId), func(t *testing.T) {
			actual := spreadsheetEncodeColumn(tc.columnId)
			if actual != tc.expected {
				t.Fatalf("expected %d, got: %d", tc.expected, actual)
			}
		})
	}
}

func BenchmarkSpreadsheetEncodeColumn(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for _, tc := range testCases {
		b.Run(fmt.Sprintf("spreadsheetEncodeColumn(%s)", tc.columnId), func(b *testing.B) {
			actual := spreadsheetEncodeColumn(tc.columnId)
			if actual != tc.expected {
				b.Fatalf("expected %d, got: %d", tc.expected, actual)
			}
		})
	}
}
