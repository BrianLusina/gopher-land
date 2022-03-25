package ledger

var successTestCases = []struct {
	name     string
	currency string
	locale   string
	entries  []Entry
	expected string
}{
	{
		name:     "empty ledger",
		currency: "USD",
		locale:   "en-US",
		entries:  nil,
		expected: `
Date       | Description               | Change
`,
	},
	{
		name:     "one entry",
		currency: "USD",
		locale:   "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-01",
				Description: "Buy present",
				Change:      -1000,
			},
		},
		expected: `
Date       | Description               | Change
01/01/2015 | Buy present               |      ($10.00)
`,
	},
	{
		name:     "credit and debit",
		currency: "USD",
		locale:   "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-02",
				Description: "Get present",
				Change:      1000,
			},
			{
				Date:        "2015-01-01",
				Description: "Buy present",
				Change:      -1000,
			},
		},
		expected: `
Date       | Description               | Change
01/01/2015 | Buy present               |      ($10.00)
01/02/2015 | Get present               |       $10.00 
`,
	},
	{
		name:     "multiple entries on same date ordered by description",
		currency: "USD",
		locale:   "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-01",
				Description: "Buy present",
				Change:      -1000,
			},
			{
				Date:        "2015-01-01",
				Description: "Get present",
				Change:      1000,
			},
		},
		expected: `
Date       | Description               | Change
01/01/2015 | Buy present               |      ($10.00)
01/01/2015 | Get present               |       $10.00 
`,
	},
	{
		name:     "final order tie breaker is change",
		currency: "USD",
		locale:   "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-01",
				Description: "Something",
				Change:      0,
			},
			{
				Date:        "2015-01-01",
				Description: "Something",
				Change:      -1,
			},
			{
				Date:        "2015-01-01",
				Description: "Something",
				Change:      1,
			},
		},
		expected: `
Date       | Description               | Change
01/01/2015 | Something                 |       ($0.01)
01/01/2015 | Something                 |        $0.00 
01/01/2015 | Something                 |        $0.01 
`,
	},
	{
		name:     "overlong descriptions",
		currency: "USD",
		locale:   "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-01",
				Description: "Freude schoner Gotterfunken",
				Change:      -123456,
			},
		},
		expected: `
Date       | Description               | Change
01/01/2015 | Freude schoner Gotterf... |   ($1,234.56)
`,
	},
	{
		name:     "euros",
		currency: "EUR",
		locale:   "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-01",
				Description: "Buy present",
				Change:      -1000,
			},
		},
		expected: `
Date       | Description               | Change
01/01/2015 | Buy present               |      (€10.00)
`,
	},
	{
		name:     "Dutch locale",
		currency: "USD",
		locale:   "nl-NL",
		entries: []Entry{
			{
				Date:        "2015-03-12",
				Description: "Buy present",
				Change:      123456,
			},
		},
		expected: `
Datum      | Omschrijving              | Verandering
12-03-2015 | Buy present               |   $ 1.234,56 
`,
	},
	{
		name:     "Dutch negative number with 3 digits before decimal point",
		currency: "USD",
		locale:   "nl-NL",
		entries: []Entry{
			{
				Date:        "2015-03-12",
				Description: "Buy present",
				Change:      -12345,
			},
		},
		expected: `
Datum      | Omschrijving              | Verandering
12-03-2015 | Buy present               |     $ 123,45-
`,
	},
	{
		name:     "American negative number with 3 digits before decimal point",
		currency: "USD",
		locale:   "en-US",
		entries: []Entry{
			{
				Date:        "2015-03-12",
				Description: "Buy present",
				Change:      -12345,
			},
		},
		expected: `
Date       | Description               | Change
03/12/2015 | Buy present               |     ($123.45)
`,
	},
}

var failureTestCases = []struct {
	name     string
	currency string
	locale   string
	entries  []Entry
}{
	{
		name:     "empty currency",
		currency: "",
		locale:   "en-US",
		entries:  nil,
	},
	{
		name:     "invalid currency",
		currency: "ABC",
		locale:   "en-US",
		entries:  nil,
	},
	{
		name:     "empty locale",
		currency: "USD",
		locale:   "",
		entries:  nil,
	},
	{
		name:     "invalid locale",
		currency: "USD",
		locale:   "nl-US",
		entries:  nil,
	},
	{
		name:     "invalid date (way too high month)",
		currency: "USD",
		locale:   "en-US",
		entries: []Entry{
			{
				Date:        "2015-131-11",
				Description: "Buy present",
				Change:      12345,
			},
		},
	},
	{
		name:     "invalid date (wrong separator)",
		currency: "USD",
		locale:   "en-US",
		entries: []Entry{
			{
				Date:        "2015-12/11",
				Description: "Buy present",
				Change:      12345,
			},
		},
	},
}
