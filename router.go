package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/scallywaag/pasta-v1/helpers"
)

type Pasta struct {
	ID        int
	Content   string `json:"content"`
	TTL       int    `json:"ttl"`
	CreatedAt string
}

func createRouter(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /pasta", func(w http.ResponseWriter, r *http.Request) {
		var pasta Pasta
		if err := json.NewDecoder(r.Body).Decode(&pasta); err != nil {
			http.Error(w, "Invalid JSON body: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		result, err := db.Exec("INSERT INTO pastas(content, ttl) VALUES(?,?)", &pasta.Content, &pasta.TTL)
		if err != nil {
			log.Fatal(err)
		}

		lastID, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		s := fmt.Sprintf("Created entity ID: %d", lastID)
		fmt.Println(s)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(s))
	})

	mux.HandleFunc("GET /pasta/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		var pasta Pasta

		err := db.QueryRow("SELECT id, content, ttl, created_at FROM pastas WHERE id = ?", id).
			Scan(&pasta.ID, &pasta.Content, &pasta.TTL, &pasta.CreatedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Pasta not found", http.StatusNotFound)
				return
			}
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		s := fmt.Sprintf("Found entity: %+v", helpers.Prettify(pasta))
		fmt.Println(s)

		w.Write([]byte(s))
	})

	return mux
}
