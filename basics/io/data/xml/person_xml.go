package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	Name string `xml:"personName"`
	Age  int    `xml:"personAge"`
}

func main() {
	b := []byte(`<Person><personName>Obama</personName><personAge>57</personAge></Person>`)
	var p Person
	// Unmarshalling
	err := xml.Unmarshal(b, &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	// Marshalling
	xmlString, _ := xml.Marshal(p)
	fmt.Printf("%s\n", xmlString)
}
