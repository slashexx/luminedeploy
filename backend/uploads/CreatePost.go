package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	// "encoding/json"
	// "log"
	"net/http"
	"scribesphere/firebaseSetup"
	"scribesphere/models"
)

func CreatePost(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()

	fmt.Println("Reached the createpost function")
	client, err := firebaseSetup.InitializeFirebase().Firestore(ctx)

	if err != nil {
		log.Fatalf("Firebase initialization error: %v", err)
		return err
	}

	defer client.Close()

	var post models.BlogPost
	json.NewDecoder(r.Body).Decode(&post)

	fmt.Println(post)

	_, err2 := client.Collection("posts").NewDoc().Create(ctx, post)

	if err2 != nil {
		fmt.Println("Error writing to firestore : ")
		log.Fatalln(err2)
	}

	return err2
}
