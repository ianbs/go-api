package main

import (
	"fmt"
	"log"
	"net/http"

	"./handlers/persons"
	"github.com/gorilla/mux"
)

func main() {
	var port = ":8000"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/persons", persons.GetAll).Methods("GET")
	router.HandleFunc("/api/persons/{id}", persons.GetOne).Methods("GET")
	fmt.Println("Server running on port: ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
