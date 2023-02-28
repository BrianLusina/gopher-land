package storage

import "io"

type MemStore struct {
	Store
}

func newMemoryStorage() Store {
	return &MemStore{}
}

func (d *MemStore) Open(name string) (io.ReadWriteCloser, error) {
	panic("not yet implemented")
}
