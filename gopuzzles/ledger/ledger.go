package ledger

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type localeInfo struct {
	currency     func(symbol string, cents int, negative bool) string
	dateFormat   string
	translations map[string]string
}

func (f localeInfo) currencyString(symbol string, cents int) string {
	negative := false
	if cents < 0 {
		cents *= -1
		negative = true
	}
	return f.currency(symbol, cents, negative)
}

func (f localeInfo) dateString(t time.Time) string {
	return t.Format(f.dateFormat)
}

var locales = map[string]localeInfo{
	"nl-NL": {
		currency:   dutchCurrencyFormat,
		dateFormat: "02-01-2006",
		translations: map[string]string{
			"date":   "Datum",
			"descr":  "Omschrijving",
			"change": "Verandering",
		},
	},
	"en-US": {
		currency:   americanCurrencyFormat,
		dateFormat: "01/02/2006",
		translations: map[string]string{
			"date":   "Date",
			"descr":  "Description",
			"change": "Change",
		},
	},
}

func dutchCurrencyFormat(symbol string, cents int, negative bool) string {
	var buf bytes.Buffer
	buf.WriteString(symbol)
	buf.WriteRune(' ')
	buf.WriteString(moneyToString(cents, ".", ","))
	if negative {
		buf.WriteRune('-')
	} else {
		buf.WriteRune(' ') // To keep alignment in report
	}
	return buf.String()
}

func americanCurrencyFormat(symbol string, cents int, negative bool) string {
	var buf bytes.Buffer
	if negative {
		buf.WriteRune('(')
	}
	buf.WriteString(symbol)
	buf.WriteString(moneyToString(cents, ",", "."))
	if negative {
		buf.WriteRune(')')
	} else {
		buf.WriteRune(' ') // To keep alignment in report
	}
	return buf.String()
}

func moneyToString(cents int, thousandsSep, decimalSep string) string {
	centsStr := fmt.Sprintf("%03d", cents) // Pad to 3 digits
	centsPart := centsStr[len(centsStr)-2:]
	rest := centsStr[:len(centsStr)-2]
	var parts []string
	for len(rest) > 3 {
		parts = append(parts, rest[len(rest)-3:])
		rest = rest[:len(rest)-3]
	}
	if len(rest) > 0 {
		parts = append(parts, rest)
	}
	revParts := make([]string, 0, len(parts))
	for i := len(parts) - 1; i >= 0; i-- {
		revParts = append(revParts, parts[i])
	}
	var buf bytes.Buffer
	buf.WriteString(strings.Join(revParts, thousandsSep))
	buf.WriteString(decimalSep)
	buf.WriteString(centsPart)
	return buf.String()
}

var currencySymbols = map[string]string{
	"USD": "$",
	"EUR": "€",
}

type entrySlice []Entry

func (e entrySlice) Len() int {
	return len(e)
}

func (e entrySlice) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e entrySlice) Less(i, j int) bool {
	if e[i].Date != e[j].Date {
		// ISO dates sort nicely
		return e[i].Date < e[j].Date
	}
	if e[i].Description != e[j].Description {
		return e[i].Description < e[j].Description
	}
	return e[i].Change < e[j].Change
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	currencySymbol, ok := currencySymbols[currency]
	if !ok {
		return "", errors.New(fmt.Sprintf("unknown currency %q", currency))
	}

	locInfo, ok := locales[locale]
	if !ok {
		return "", errors.New(fmt.Sprintf("unknown locale %q", locale))
	}

	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)

	sort.Sort(entrySlice(entriesCopy))
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%-10s | %-25s | %s\n",
		locInfo.translations["date"],
		locInfo.translations["descr"],
		locInfo.translations["change"]),
	)

	for _, entry := range entriesCopy {
		date, err := time.Parse("2006-01-02", entry.Date)
		if err != nil {
			return "", err
		}
		description := entry.Description
		if len(description) > 25 {
			description = description[:22] + "..."
		}
		buf.WriteString(fmt.Sprintf("%-10s | %-25s | %13s\n",
			locInfo.dateString(date),
			description,
			locInfo.currencyString(currencySymbol, entry.Change)))
	}
	return buf.String(), nil
}
