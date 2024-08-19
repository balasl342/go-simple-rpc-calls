package rpc

import (
	"log"
	"net"
	"net/rpc"

	"github.com/balasl342/go-simple-rpc-calls/internal/shared"
)

func StartServer() {
	serviceObj := new(shared.Service)
	rpc.Register(serviceObj)
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
