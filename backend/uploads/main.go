package main

import (
	"fmt"
	"net/http"
	"scribesphere/handlers"
	"scribesphere/models"
	"scribesphere/cors"
)

var posts []models.BlogPost

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello, check out URLs: /posts for all posts and /post/{id} for specific post")
	})

	// Combine the GET and POST for /posts into a single handler
	http.Handle("/posts", cors.EnableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetAllPosts(w, r) // Call the handler to get all posts
		case http.MethodPost:
			handlers.CreatePost(w, r) // Call the handler to create a post
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	// Handle specific post actions
	http.Handle("/post/", cors.EnableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetPost(w, r) // GET a specific post
		case http.MethodPut:
			handlers.UpdatePost(w, r) // PUT to update a post
		case http.MethodDelete:
			handlers.DeletePost(w, r) // DELETE a post
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	fmt.Println("The server is now listening at port : 8080 (http://localhost:8080)")
	http.ListenAndServe(":8080", nil)
}
