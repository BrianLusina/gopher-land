package storage

import "io"

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

func NewStore(st StorageType) Store {
	switch st {
	case DiskStorage:
		return newDiskStorage()
	case MemoryStorage:
		return newMemoryStorage()
	case TempStorage:
		return newTempStorage()
	default:
		return nil
	}
}
