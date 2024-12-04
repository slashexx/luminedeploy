package handlers

import (
	"net/http"
)

func DeletePost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/post/"):]

	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
		}
	}

	http.Error(w, "Post to be deleted not found", http.StatusNotFound)
}
