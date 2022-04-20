package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Starting the listener...")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			return
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return // terminate program
		}
		fmt.Printf("Received data: %v", string(buf))
	}
}
