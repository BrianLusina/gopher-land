package readerwriter

import (
	"strings"
	"testing"
	"testing/iotest"
)

func TestFoo(t *testing.T) {
	err := foo(iotest.TimeoutReader(
		strings.NewReader("abcDEFghi")),
	)
	if err != nil {
		t.Fatal(err)
	}
}
