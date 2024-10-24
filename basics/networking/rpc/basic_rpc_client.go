package main

import (
	"fmt"
	"gopherland/basics/networking/rpc/rpcobjects"
	"log"
	"net/rpc"
)

const serverAddress = "localhost"

func run_basic_rpc() {
	client, err := rpc.DialHTTP("tcp", serverAddress+":3001")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}
	// Synchronous call
	args := &rpcobjects.Args{N: 7, M: 8}
	var reply int
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Args error:", err)
	}
	fmt.Printf("Args: %d * %d = %d", args.N, args.M, reply)
}
