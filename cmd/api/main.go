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

type APIInfo struct {
	Name     string `json:"name"`
	Language string `json:"language"`
	Category string `json:"category"`
}

type Meta struct {
	API   APIInfo `json:"api"`
	Count int     `json:"count,omitempty"`
}

type PostsResponse struct {
	Meta Meta `json:"meta"`
	Data struct {
		Posts []Post `json:"posts"`
	} `json:"data"`
}

type PostResponse struct {
	Meta Meta `json:"meta"`
	Data struct {
		Post Post `json:"post"`
	} `json:"data"`
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

func writeJSONError(w http.ResponseWriter, status int, code string, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := ErrorResponse{
		Error: ErrorDetail{
			Code:    code,
			Message: message,
		},
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("failed to encode error response: %v", err)
	}
}

func requireAPIKey(apiKey string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-API-Key") != apiKey {
			writeJSONError(
				w,
				http.StatusUnauthorized,
				"UNAUTHORIZED",
				"Invalid API key",
			)
			return
		}

		next(w, r)
	}
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

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY is required")
	}

	apiInfo := APIInfo{
		Name:     "go-api",
		Language: "Go",
		Category: "public",
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	mux.HandleFunc("GET /api/v1/posts", requireAPIKey(apiKey, func(w http.ResponseWriter, r *http.Request) {
		result := posts

		username := r.URL.Query().Get("username")
		if username != "" {
			result = []Post{}

			for _, post := range posts {
				if post.User.Username == username {
					result = append(result, post)
				}
			}
		}

		response := PostsResponse{
			Meta: Meta{
				API:   apiInfo,
				Count: len(result),
			},
		}
		response.Data.Posts = result

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("failed to encode posts: %v", err)
		}
	}))

	mux.HandleFunc("GET /api/v1/posts/{id}", requireAPIKey(apiKey, func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			writeJSONError(
				w,
				http.StatusBadRequest,
				"INVALID_POST_ID",
				"Invalid post ID",
			)
			return
		}

		for _, post := range posts {
			if post.ID == id {
				response := PostResponse{
					Meta: Meta{
						API: apiInfo,
					},
				}
				response.Data.Post = post

				w.Header().Set("Content-Type", "application/json")

				if err := json.NewEncoder(w).Encode(response); err != nil {
					log.Printf("failed to encode post: %v", err)
				}

				return
			}
		}

		writeJSONError(
			w,
			http.StatusNotFound,
			"NOT_FOUND",
			"Post not found",
		)
	}))

	addr := ":3000"

	log.Printf("API server listening on %s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
