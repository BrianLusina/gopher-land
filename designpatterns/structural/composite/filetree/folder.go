package filetree

import "fmt"

type Folder struct {
	components []Component
	name       string
}

func (folder *Folder) search(keyword string) {
	fmt.Printf("Searching through folder %s for keyword %s", folder.name, keyword)
	for _, c := range folder.components {
		c.search(keyword)
	}
}

func (folder *Folder) GetName() string {
	return folder.name
}

func (folder *Folder) add(c Component) {
	folder.components = append(folder.components, c)
}
