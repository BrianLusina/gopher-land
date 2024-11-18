package main

import (
	"fmt"
	"os"
	"strings"
)

func greeting() {
	who := "Me "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], ", ")
	}
	fmt.Println("Hello", who)
}
