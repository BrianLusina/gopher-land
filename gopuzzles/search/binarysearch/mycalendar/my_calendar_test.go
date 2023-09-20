package mycalendar

import (
	"fmt"
	"testing"
)

type testCase struct {
	start    int
	end      int
	expected bool
}

var testCases = []testCase{
	{
		start:    10,
		end:      20,
		expected: true,
	},
	{
		start:    15,
		end:      25,
		expected: false,
	},
	{
		start:    20,
		end:      30,
		expected: true,
	},
}

func TestMyCalendar(t *testing.T) {
	myCalendar := NewMyCalendar()

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("myCalendar.Book(start=%d, end=%d)", tc.start, tc.end), func(t *testing.T) {
			actual := myCalendar.Book(tc.start, tc.end)
			if actual != tc.expected {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func BenchmarkMyCalendar(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		myCalendar := NewMyCalendar()

		for _, tc := range testCases {
			b.Run(fmt.Sprintf("myCalendar.Book(start=%d, end=%d)", tc.start, tc.end), func(b *testing.B) {
				actual := myCalendar.Book(tc.start, tc.end)
				if actual != tc.expected {
					b.Fatalf("expected %v, got %v", tc.expected, actual)
				}
			})
		}
	}
}
