package peeking

import (
	"testing"
)

type peekingIteratorTestCase struct {
	description string
	items       []int
	firstNext   int
	firstPeek   int
	secondNext  int
	thirdNext   int
}

var peekingIteratorTestCases = []peekingIteratorTestCase{
	{
		description: "PeekingIterator -> next -> peek -> next -> next -> hasNext",
		items:       []int{1, 2, 3},
		firstNext:   1,
		firstPeek:   2,
		secondNext:  2,
		thirdNext:   3,
	},
}

func TestPeekingIterator(t *testing.T) {
	for _, tc := range peekingIteratorTestCases {
		t.Run(tc.description, func(t *testing.T) {
			peekingIterator := NewPeekingIterator(tc.items)

			firstNext := peekingIterator.next()
			if firstNext != tc.firstNext {
				t.Fatalf("PeekingIterator(%v).next() = %d, expected = %d", tc.items, firstNext, tc.firstNext)
			}

			actualPeek := peekingIterator.peek()
			if actualPeek != tc.firstPeek {
				t.Fatalf("PeekingIterator(%v).peek() = %d, expected = %d", tc.items, actualPeek, tc.firstPeek)
			}

			if next := peekingIterator.next(); next != tc.secondNext {
				t.Fatalf("PeekingIterator(%v).next() = %d, expected = %d", tc.items, next, tc.secondNext)
			}

			if next := peekingIterator.next(); next != tc.thirdNext {
				t.Fatalf("PeekingIterator(%v).next() = %d, expected = %d", tc.items, next, tc.thirdNext)
			}

			if hasNext := peekingIterator.hasNext(); hasNext != false {
				t.Fatalf("PeekingIterator(%v).hasNext() = %v, expected = %v", tc.items, hasNext, false)
			}
		})
	}
}
