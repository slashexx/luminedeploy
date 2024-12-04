package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Struct to hold user input for Dockerfile generation
type DockerfileConfig struct {
	BaseImage        string `json:"baseImage"`
	WorkingDirectory string `json:"workingDirectory"`
	CopyCommand      string `json:"copyCommand"`
	InstallCommand   string `json:"installCommand"`
	StartCommand     string `json:"startCommand"`
}

// GenerateGoDockerfile handles generating the Dockerfile based on user input
func GenerateGoDockerfile(w http.ResponseWriter, r *http.Request) {
	var config DockerfileConfig

	// Parse the JSON body
	err := json.NewDecoder(r.Body).Decode(&config)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid input: %v", err), http.StatusBadRequest)
		return
	}

	// Default values if not provided by the user
	if config.BaseImage == "" {
		config.BaseImage = "golang:1.22"
	}
	if config.WorkingDirectory == "" {
		config.WorkingDirectory = "/app"
	}
	if config.CopyCommand == "" {
		config.CopyCommand = "COPY . ."
	}
	if config.InstallCommand == "" {
		config.InstallCommand = "RUN go mod tidy && go build -o app"
	}
	if config.StartCommand == "" {
		config.StartCommand = "CMD [\"./app\"]"
	}

	// Dockerfile content using the user-defined or default values
	dockerfileContent := fmt.Sprintf(`# Stage 1: Build
FROM %s AS builder
WORKDIR %s
%s
%s

# Stage 2: Run
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder %s/app .
%s
`, config.BaseImage, config.WorkingDirectory, config.CopyCommand, config.InstallCommand, config.WorkingDirectory, config.StartCommand)

	// Default project directory
	projectPath := "./aiseHi" // Modify this path as needed

	// Ensure the directory exists
	err = os.MkdirAll(projectPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create directory %s: %v", projectPath, err)
	}

	// Write the Dockerfile
	dockerfilePath := filepath.Join(projectPath, "Dockerfile")
	file, err := os.Create(dockerfilePath)
	if err != nil {
		log.Fatalf("Failed to create Dockerfile: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(dockerfileContent)
	if err != nil {
		log.Fatalf("Failed to write Dockerfile content: %v", err)
	}

	log.Println("Dockerfile successfully created in:", projectPath)

	// Respond to the client that the Dockerfile has been created
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Dockerfile successfully created",
		"path":    dockerfilePath,
	})
}
