package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type MyAddress struct {
	Type    string
	City    string
	Country string
}

type Card struct {
	FirstName string
	LastName  string
	Addresses []*MyAddress
	Remark    string
}

func gob_read_file() {
	file, _ := os.OpenFile("./output/vcard.gob", os.O_RDONLY, 0666)

	reader := bufio.NewReader(file)
	decoder := gob.NewDecoder(reader)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Error closing file:", err)
		}
	}(file)

	var vcard Card
	err := decoder.Decode(&vcard)
	if err != nil {
		log.Fatalf("Error in decoding gob: %s", err)
	}
	fmt.Printf("%q\n", vcard)
}
