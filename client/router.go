package main

import (
	//Standard Libs
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"strconv"

	//Custom Libs
	"github.com/gorilla/mux"
)

type Product struct {
	Id     int    `gorm:"primary_key;autoIncrement" json:"id"`
	Name   string `gorm:"column:product_name" json:"product_name"`
	ProdID int    `json:"product_id" gorm:"column:product_id"`
	Price  int    `json:"product_price" gorm:"column:product_price"`
}

type Service struct{}

// Create product
func (s Service) CreateProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("enter")
	client, err := rpc.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("Error connecting to RPC server:", err)
	}
	var req Product
	json.NewDecoder(r.Body).Decode(&req)
	reqs := Product{
		Name:   req.Name,
		ProdID: req.ProdID,
		Price:  req.Price,
	}
	var reply Product
	err = client.Call("Service.CreateProducts", reqs, &reply)
	if err != nil {
		fmt.Println("Error calling Concatenate:", err)
	}
	w.Header().Set("Content-TYpe", "application/json")
	json.NewEncoder(w).Encode(reply)
}

// Get products
func (s Service) GetProductsByID(w http.ResponseWriter, r *http.Request) {
	client, err := rpc.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("Error connecting to RPC server:", err)
	}
	var reply Product
	vars := mux.Vars(r)
	id := vars["id"]
	pid, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	reqs := Product{
		Id: pid,
	}
	err = client.Call("Service.GetProductsByID", reqs, &reply)
	if err != nil {
		fmt.Println("Error calling Concatenate:", err)
	}
	w.Header().Set("Content-TYpe", "application/json")
	json.NewEncoder(w).Encode(reply)
}
