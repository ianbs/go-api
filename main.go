package main

import (
	"log"
	"net/http"

	"./handlers/hellodao"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api", hellodao.HelloGet).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
