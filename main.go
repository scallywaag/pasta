package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := createRouter()
	fmt.Println("Starting server on localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
