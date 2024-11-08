package main

import (
	"log"
	"net/http"

	"github.com/train-do/Golang-Restfull-API/router"
)

func main() {
	r := router.NewRouter()

	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
