package main

import (
	"log"
	"net"
	"net/rpc"
)

func main() {
	service1 := new(Service)
	rpc.Register(service1)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("Error starting RPC server:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting connection:", err)
		}
		go rpc.ServeConn(conn)
	}
}
