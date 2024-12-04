package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "os"
	// "path/filepath"
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

	// Generate the Dockerfile content
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

	// Set response headers
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", `attachment; filename="Dockerfile"`)

	// Write the Dockerfile content to the response
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(dockerfileContent))
	if err != nil {
		log.Printf("Error writing Dockerfile content: %v", err)
		http.Error(w, "Failed to generate Dockerfile", http.StatusInternalServerError)
		return
	}
}
