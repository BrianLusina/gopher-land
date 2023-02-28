package storage

import "io"

type TempStore struct {
	Store
}

func newTempStorage() Store {
	return &TempStore{}
}

func (d *TempStore) Open(name string) (io.ReadWriteCloser, error) {
	panic("not yet implemented")
}
