package main

import (
	"fmt"
	"net/http"
)

func createRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/path", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "got path\n")
	})

	mux.HandleFunc("/task/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "handling task with id=%v\n", id)
	})

	return mux
}
