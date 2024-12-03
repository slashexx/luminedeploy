package main

import (
	"log"
	"net/http"
	// "time"

	"github.com/gorilla/mux"
	"lumine/backend/routes"
)

func main() {
	// Initialize the Gorilla Mux router
	r := mux.NewRouter()

	// Register API routes
	routes.RegisterRoutes(r)

	// Set up a server with custom configurations
	server := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		// WriteTimeout: 15 * time.Second,
		// ReadTimeout:  15 * time.Second,
		// IdleTimeout:  60 * time.Second,
	}

	log.Println("Server is running on http://localhost:8080")
	// Start the server
	log.Fatal(server.ListenAndServe())
}
