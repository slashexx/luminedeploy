package handlers

import (
	"encoding/json"
	"net/http"
	"scribesphere/models"
)

var posts []models.BlogPost

// GetAllPosts handles the GET request to retrieve all blog posts.
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if len(posts) == 0 {
		w.WriteHeader(http.StatusNotFound) // 404 Not Found
		json.NewEncoder(w).Encode(map[string]string{"message": "No posts found"})
		return
	}

	w.WriteHeader(http.StatusOK) // 200 OK
	json.NewEncoder(w).Encode(posts)
}
