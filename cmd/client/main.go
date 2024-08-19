package main

import (
	"net/http"

	"github.com/balasl342/go-simple-rpc-calls/internal/client"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	var serv client.ServiceWrapper
	router.HandleFunc("/create", serv.CreateProducts).Methods("POST")
	router.HandleFunc("/get/{id}", serv.GetProductsByID).Methods("GET")

	http.ListenAndServe(":8080", router)

}
