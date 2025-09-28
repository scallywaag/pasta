package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/scallywaag/pasta-v1/helpers"
)

type Pasta struct {
	Content string `json:"content"`
	Ttl     int    `json:"ttl"`
}

func createRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /pasta", func(w http.ResponseWriter, r *http.Request) {
		var pasta Pasta

		if err := json.NewDecoder(r.Body).Decode(&pasta); err != nil {
			http.Error(w, "Invalid JSON body: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		helpers.PrintJson(pasta)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("POST received"))
	})

	mux.HandleFunc("GET /pasta/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "getting pasta with id=%v\n", id)
	})

	return mux
}
