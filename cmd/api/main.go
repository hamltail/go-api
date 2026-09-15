package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type User struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
}

type Post struct {
	ID        int    `json:"id"`
	User      User   `json:"user"`
	Content   string `json:"content"`
	PostedOn  string `json:"postedOn"`
	CreatedAt string `json:"createdAt"`
}

func main() {
	data, err := os.ReadFile("data/posts.json")
	if err != nil {
		log.Fatal(err)
	}

	var posts []Post

	if err := json.Unmarshal(data, &posts); err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded %d posts", len(posts))

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	mux.HandleFunc("GET /api/v1/posts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(posts); err != nil {
			log.Printf("failed to encode posts: %v", err)
		}
	})

	addr := ":3000"

	log.Printf("API server listening on %s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
