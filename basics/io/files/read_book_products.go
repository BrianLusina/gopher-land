package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type (
	Book struct {
		title    string
		price    float64
		quantity int
	}

	books []Book
)

func readFile(filename string) books {
	result := make(books, 0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if len(record) == 3 {
			title, p, qty := record[0], record[1], record[2]
			quantity, err := strconv.Atoi(qty)

			if err != nil {
				_ = fmt.Sprintf("%s is not a valid quantity for record %s", qty, record)
				continue
			}

			price, err := strconv.ParseFloat(p, 64)
			if err != nil {
				_ = fmt.Sprintf("%s is not a valid price for record %s", p, record)
				continue
			}

			book := Book{title, price, quantity}
			result = append(result, book)
		}
	}
	return result
}

func main() {
	var filename string
	fmt.Println("Please enter the name of the file to read:")
	_, err := fmt.Scanln(&filename)
	if err != nil {
		log.Fatal(err)
	}

	books := readFile(filename)

	for _, book := range books {
		fmt.Printf("%s: %.2f %d \n", book.title, book.price, book.quantity)
	}
}
