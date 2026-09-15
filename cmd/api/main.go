package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
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

	mux.HandleFunc("GET /api/v1/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		for _, post := range posts {
			if post.ID == id {
				w.Header().Set("Content-Type", "application/json")

				if err := json.NewEncoder(w).Encode(post); err != nil {
					log.Printf("failed to encode post: %v", err)
				}

				return
			}
		}

		http.Error(w, "Post not found", http.StatusNotFound)
	})

	addr := ":3000"

	log.Printf("API server listening on %s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
