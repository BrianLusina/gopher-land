package readerwriter

import (
	"fmt"
	"io"
)

func foo(r io.Reader) error {
	b, err := readAll(r, 3)
	if err != nil {
		return err
	}

	fmt.Printf("read %d bytes, %v", len(b), b)

	return nil
}

func readAll(r io.Reader, retries int) ([]byte, error) {
	b := make([]byte, 0, 512)
	for {
		if len(b) == cap(b) {
			b = append(b, 0)[:len(b)]
		}
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == io.EOF {
				return b, nil
			}
			retries--
			// tolerates retries
			if retries < 0 {
				return b, err
			}
		}
	}
}
