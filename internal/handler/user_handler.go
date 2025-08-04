package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"go-cassandra-demo-service/internal/model"
	"go-cassandra-demo-service/internal/service"

	"github.com/gocql/gocql"
)

func MakeUserHandler(session *gocql.Session) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("/user received")
		if r.Method != http.MethodPost {
			log.Print("Method not allowed")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := logBody(w, r); err != "" {
			log.Print(err)
		}

		var u model.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		u.ID = gocql.TimeUUID()

		if err := service.InsertUser(session, &u); err != nil {
			log.Printf("Insert error: %v", err)
			http.Error(w, "Failed to insert user", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(u)
	}
}

func logBody(w http.ResponseWriter, r *http.Request) string {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusBadRequest)
		return "Error reading body"
	}

	// Log the raw body
	log.Printf("Request Body: %s", string(bodyBytes))

	// Restore the io.ReadCloser so we can decode it again
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return ""
}
