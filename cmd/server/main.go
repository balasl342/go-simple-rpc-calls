package main

import (
	"log"

	"github.com/balasl342/go-simple-rpc-calls/internal/server/db"
	"github.com/balasl342/go-simple-rpc-calls/internal/server/rpc"
)

func init() {
	err := db.DBconnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

func main() {
	// Start RPC server
	rpc.StartServer()
}
