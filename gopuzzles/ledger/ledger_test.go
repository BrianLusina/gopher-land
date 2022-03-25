package ledger

import (
	"reflect"
	"strings"
	"testing"
)

func TestFormatLedgerSuccess(t *testing.T) {
	for _, tt := range successTestCases {
		actual, err := FormatLedger(tt.currency, tt.locale, tt.entries)
		// We don't expect errors for any of the test cases
		if err != nil {
			var _ error = err
			t.Fatalf("FormatLedger for input named %q returned error %q. Error not expected.",
				tt.name, err)
		}
		expected := tt.expected[1:] // Strip initial newline
		if actual != expected {
			t.Fatalf("FormatLedger for input named %q was expected to return...\n%s\n...but returned...\n%s",
				tt.name, strings.ReplaceAll(expected, " ", "_"), strings.ReplaceAll(actual, " ", "_"))
		}
	}
}

func TestFormatLedgerFailure(t *testing.T) {
	for _, tt := range failureTestCases {
		_, err := FormatLedger(tt.currency, tt.locale, tt.entries)
		if err == nil {
			t.Fatalf("FormatLedger for input %q should have failed but didn't.", tt.name)
		}
	}
}

func TestFormatLedgerNotChangeInput(t *testing.T) {
	entries := []Entry{
		{
			Date:        "2015-01-02",
			Description: "Freude schöner Götterfunken",
			Change:      1000,
		},
		{
			Date:        "2015-01-01",
			Description: "Buy present",
			Change:      -1000,
		},
	}
	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)
	FormatLedger("USD", "en-US", entries)
	if !reflect.DeepEqual(entries, entriesCopy) {
		t.Fatalf("FormatLedger modifies the input entries array")
	}
}

func BenchmarkFormatLedger(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range successTestCases {
			FormatLedger(tt.currency, tt.locale, tt.entries)
		}
	}
}
