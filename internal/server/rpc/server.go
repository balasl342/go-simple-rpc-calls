package rpc

import (
	"log"
	"net"
	"net/rpc"

	"github.com/balasl342/go-simple-rpc-calls/internal/server/db"
	"github.com/balasl342/go-simple-rpc-calls/internal/server/repository"
)

func StartServer() {
	productRepo := &repository.ProductRepository{Db: db.Config.DB}
	rpcService := &RPCService{Repo: productRepo}
	rpc.Register(rpcService)
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
