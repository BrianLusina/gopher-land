package circularlinked

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseAppend[T any] struct {
	data       []T
	appendData T
	expected   []T
}

var testCasesAppend = []testCaseAppend[any]{
	{
		data:       []any{1, 2, 3, 4, 5, 6},
		appendData: 7,
		expected:   []any{1, 2, 3, 4, 5, 6, 7},
	},
}

// TestAppend tests adding a new data item to a circular linked list
func TestAppend(t *testing.T) {
	for _, tc := range testCasesAppend {
		testName := fmt.Sprintf("Initial CircularLinkedList of data={%v} should append(%v) to get expected={%v}", tc.data, tc.appendData, tc.expected)
		t.Run(testName, func(t *testing.T) {
			linkedList := New[any]()
			for _, data := range tc.data {
				linkedList.Append(data)
			}
			actualHead := linkedList.HeadNode()
			assert.NotNil(t, actualHead)
			assert.Equal(t, tc.data[0], actualHead.Data)

			// append new data
			linkedList.Append(tc.appendData)

			actualHead2 := linkedList.HeadNode()
			assert.Equal(t, tc.data[0], actualHead2.Data)

			actualData := []any{}
			current := actualHead2

			for current != nil {
				actualData = append(actualData, current.Data)

				if current.Next == actualHead2 {
					break
				}
				current = current.Next
			}

			assert.ElementsMatch(t, tc.expected, actualData)
		})
	}
}

func BenchmarkAppend(b *testing.B) {
	if testing.Short() {
		b.Skipf("Skipping benchmark tests in short mode")
	}

	for _, tc := range testCasesAppend {
		testName := fmt.Sprintf("Initial CircularLinkedList of data={%v} should append(%v) to get expected={%v}", tc.data, tc.appendData, tc.expected)
		b.Run(testName, func(t *testing.B) {
			linkedList := New[any]()
			for _, data := range tc.data {
				linkedList.Append(data)
			}
			actualHead := linkedList.HeadNode()
			assert.NotNil(t, actualHead)
			assert.Equal(t, tc.data[0], actualHead.Data)

			// append new data
			linkedList.Append(tc.appendData)

			actualHead2 := linkedList.HeadNode()
			assert.Equal(t, tc.data[0], actualHead2.Data)

			actualData := []any{}
			current := actualHead2

			for current != nil {
				actualData = append(actualData, current.Data)

				if current.Next == actualHead2 {
					break
				}
				current = current.Next
			}

			assert.ElementsMatch(t, tc.expected, actualData)
		})
	}
}
