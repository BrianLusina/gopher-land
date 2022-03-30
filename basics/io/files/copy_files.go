package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func CopyFile(dst, src string) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}

	defer func(srcFile *os.File) {
		err := srcFile.Close()
		if err != nil {
			panic(err)
		}
	}(srcFile)

	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}

	defer func(dstFile *os.File) {
		err := dstFile.Close()
		if err != nil {
			panic(err)
		}
	}(dstFile)

	return io.Copy(dstFile, srcFile)
}

func main() {
	var sourceFile string
	var destinationFile string
	fmt.Println("Please enter the source file to read:")
	_, err := fmt.Scanln(&sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Please enter the destination file to read:")
	_, err = fmt.Scanln(&destinationFile)
	if err != nil {
		log.Fatal(err)
	}

	_, err = CopyFile(destinationFile, sourceFile)
	if err != nil {
		return // handle error
	}
	fmt.Println("Copy done!")
}
