package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	var handler Service
	router.HandleFunc("/create", handler.CreateProducts).Methods("POST")
	router.HandleFunc("/get/{id}", handler.GetProductsByID).Methods("GET")

	http.ListenAndServe(":8080", router)

}
