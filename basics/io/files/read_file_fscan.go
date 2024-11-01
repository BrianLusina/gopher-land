package main

import (
	"fmt"
	"os"
)

func read_file_fscan() {
	file, err := os.Open("./files/products2.txt")
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	var (
		col1, col2, col3 []string
	)

	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3) //scans until newline
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}
