package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "pasta.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := createRouter(db)
	fmt.Println("Starting server on localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
