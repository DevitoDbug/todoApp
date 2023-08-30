package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	port := ":8080"

	r := mux.NewRouter()

	http.Handle("/", r)

	log.Printf("Starting server at port%v \n", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalf("Could not start sever:\n %v \n", err)
		return
	}
}
