package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type (
	Book2 struct {
		title    string
		price    float64
		quantity int
	}

	books2 []Book2
)

// readFile2 is an alternative to readFile.
func readFile2(filename string) books2 {
	bks := make(books2, 1)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error %s opening file products.txt: ", err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		// read one line from the file:
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		// remove \r and \n so 2 in Windows, in Linux only \n, so 1:
		line = string(line[:len(line)-2])
		//fmt.Printf("The input was: -%s-", line)

		strSl := strings.Split(line, ";")
		book := new(Book2)
		book.title = strSl[0]
		book.price, err = strconv.ParseFloat(strSl[1], 32)
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}
		//fmt.Printf("The quan was:-%s-", strSl[2])
		book.quantity, err = strconv.Atoi(strSl[2])
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}
		if bks[0].title == "" {
			bks[0] = *book
		} else {
			bks = append(bks, *book)
		}
	}
	fmt.Println("We have read the following books from the file: ")
	return bks
}

func read_book_products_2() {
	var filename string
	fmt.Println("Please enter the name of the file to read:")
	_, err := fmt.Scanln(&filename)
	if err != nil {
		log.Fatal(err)
	}

	books := readFile2(filename)

	for _, book := range books {
		fmt.Printf("%s: %.2f %d \n", book.title, book.price, book.quantity)
	}
}
