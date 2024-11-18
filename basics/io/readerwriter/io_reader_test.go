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

func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo(iotest.TimeoutReader(
			strings.NewReader("abcDEFghi")),
		)
	}
}
