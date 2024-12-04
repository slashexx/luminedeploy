package handlers

import (
	"encoding/json"
	"net/http"
	// "scribesphere/models"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/post/"):]
	for i := 0; i < len(posts); i++ {
		if posts[i].ID == id {
			json.NewEncoder(w).Encode(posts[i])
			return
		}
	}

	http.Error(w, "Specified post does not exist", http.StatusNotFound)
}
