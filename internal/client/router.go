package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"strconv"

	"github.com/balasl342/go-simple-rpc-calls/internal/shared"

	"github.com/gorilla/mux"
)

type ServiceWrapper struct{}

// Create product
func (*ServiceWrapper) CreateProducts(w http.ResponseWriter, r *http.Request) {
	client, err := rpc.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("Error connecting to RPC server:", err)
	}
	var req shared.Product
	json.NewDecoder(r.Body).Decode(&req)
	reqs := shared.Product{
		Name:   req.Name,
		ProdID: req.ProdID,
		Price:  req.Price,
	}
	var reply shared.Product
	err = client.Call("RPCService.CreateProducts", reqs, &reply)
	if err != nil {
		fmt.Println("Error calling Concatenate:", err)
	}
	w.Header().Set("Content-TYpe", "application/json")
	json.NewEncoder(w).Encode(reply)
}

// Get products
func (*ServiceWrapper) GetProductsByID(w http.ResponseWriter, r *http.Request) {
	client, err := rpc.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("Error connecting to RPC server:", err)
	}
	var reply shared.Product
	vars := mux.Vars(r)
	id := vars["id"]
	pid, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	reqs := shared.Product{
		Id: pid,
	}
	err = client.Call("RPCService.GetProductsByID", reqs, &reply)
	if err != nil {
		fmt.Println("Error calling Concatenate:", err)
	}
	w.Header().Set("Content-TYpe", "application/json")
	json.NewEncoder(w).Encode(reply)
}
