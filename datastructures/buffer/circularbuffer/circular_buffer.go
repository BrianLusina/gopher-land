package circularbuffer

import "gopherland/datastructures/buffer"

// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

type Buffer struct {
	items []byte
	// w next position to write
	w int
	// r next position to read
	r    int
	size int
}

func NewBuffer(size int) *Buffer {
	if size <= 0 {
		panic("buffer size must be > 0")
	}

	return &Buffer{
		size:  size,
		w:     0,
		r:     0,
		items: make([]byte, size),
	}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.r >= b.w {
		return 0, buffer.ErrEmptyBuffer
	}
	item := b.items[b.r%b.size]
	b.r++
	return item, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.w-b.r == b.size {
		return buffer.ErrFullBuffer
	}

	b.items[b.w%b.size] = c
	b.w++

	return nil
}

func (b *Buffer) Overwrite(c byte) {
	b.items[b.w%b.size] = c

	if b.w-b.r == b.size {
		b.r++
	}
	b.w++
}

func (b *Buffer) Reset() {
	b.w = 0
	b.r = 0
}
