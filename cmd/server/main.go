package main

import (
	"log"

	"github.com/balasl342/go-simple-rpc-calls/internal/server/db"
	"github.com/balasl342/go-simple-rpc-calls/internal/server/rpc"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dbErr := db.DBconnection()
	if dbErr != nil {
		log.Fatalf("Failed to connect to database: %v", dbErr)
	}
}

func main() {
	// Start RPC server
	rpc.StartServer()
}
