package filetree

import "fmt"

type File struct {
	name string
}

func (file *File) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s", keyword, file.name)
}

func (file *File) GetName() string {
	return file.name
}
