package main

import (
	"flag"
	"fmt"
)

func flags() {
	strPtr := flag.String("lang", "Go", "a string")
	numPtr := flag.Int("num", 108, "an int")
	boolPtr := flag.Bool("truth", false, "a bool")
	var str string
	flag.StringVar(&str, "str", "Crystal", "a string variable")
	flag.Parse()
	fmt.Println("lang:", *strPtr)
	fmt.Println("num:", *numPtr)
	fmt.Println("truth:", *boolPtr)
	fmt.Println("str:", str)
	fmt.Println("tail:", flag.Args())
}
