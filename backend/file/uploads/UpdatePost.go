package handlers

import (
	"encoding/json"
	"net/http"
	"scribesphere/models"
)

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/post/"):]

	var updatedPost models.BlogPost
	json.NewDecoder(r.Body).Decode(&updatedPost)

	for i, post := range posts {
		if post.ID == id {
			posts[i] = updatedPost
			json.NewEncoder(w).Encode(posts[i])
			return
		}
	}

	http.Error(w, "Post not found", http.StatusNotFound)
}
