package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	js, err := json.Marshal(vc)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("JSON format: %s\n", js)

	// using an encoder:
	file, _ := os.OpenFile("./output/vcard.json", os.O_CREATE|os.O_WRONLY, 0)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
			log.Println("Error in closing file")
		}
	}(file)

	enc := json.NewEncoder(file)
	err = enc.Encode(vc)

	if err != nil {
		log.Println(err)
		log.Println("Error in encoding json")
	}
}
