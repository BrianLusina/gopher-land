package storage

import "io"

type DiskStore struct {
	Store
}

func newDiskStorage() Store {
	return &DiskStore{}
}

func (d *DiskStore) Open(name string) (io.ReadWriteCloser, error) {
	panic("not yet implemented")
}
