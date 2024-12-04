package main

import (
	"log"
	"net/http"
	"time"

	"lumine/backend/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the Gorilla Mux router
	r := mux.NewRouter()

	// Register API routes
	routes.RegisterRoutes(r)

	// Enable CORS to allow requests from the frontend (e.g., http://localhost:3000)
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	// Set up a server with custom configurations
	server := &http.Server{
		Handler:      handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r), // Apply CORS middleware
		Addr:         ":8080",                                                          // Server listens on port 8080
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Server is running on http://localhost:8080")
	// Start the server
	log.Fatal(server.ListenAndServe())
}
