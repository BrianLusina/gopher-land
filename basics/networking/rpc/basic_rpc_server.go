package main

import (
	"fmt"
	"gopherland/basics/networking/rpc/rpcobjects"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func main() {
	calc := new(rpcobjects.Args)
	err := rpc.Register(calc)
	if err != nil {
		fmt.Println("There was an error:", err.Error())
	}
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", "0.0.0.0:3001")
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}
	go func() {
		err := http.Serve(listener, nil)
		if err != nil {
			log.Fatal("Starting RPC-server -serve error:", err)
		}
	}()
	time.Sleep(1000e9)
}
