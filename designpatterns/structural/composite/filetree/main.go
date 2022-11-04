package filetree

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}
	file4 := &File{name: "File4"}

	folder1 := &Folder{name: "Folder1"}
	folder2 := &Folder{name: "Folder2"}
	folder3 := &Folder{name: "Folder3"}
	folder4 := &Folder{name: "Folder4"}

	folder1.add(file1)

	folder2.add(file2)
	folder3.add(folder1)
	folder4.add(file3)
	folder4.add(folder3)
	folder4.add(file4)

	folder4.search("rose")
}
