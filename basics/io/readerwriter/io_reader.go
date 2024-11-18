package readerwriter

import (
	"fmt"
	ioutils "gopherland/pkg/utils/io"
	"io"
)

func foo(r io.Reader) error {
	b, err := ioutils.ReadAllWithRetries(r, 3)
	if err != nil {
		return err
	}

	fmt.Printf("read %d bytes, %v", len(b), b)

	return nil
}
