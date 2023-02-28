package storage

import "fmt"

func client() {
	s := NewStore(MemoryStorage)
	f, err := s.Open("file")
	if err != nil {
		panic(err)
	}

	n, _ := f.Write([]byte("data"))
	defer f.Close()

	fmt.Sprintf("Lines written %d", n)
}
