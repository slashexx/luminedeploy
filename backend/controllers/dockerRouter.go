package controllers

import (
	"log"
	"os"
	"path/filepath"
	"net/http"
)

func GenerateGoDockerfile(w http.ResponseWriter, r *http.Request) {
	// Dockerfile content
	dockerfileContent := `# Stage 1: Build
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o app

# Stage 2: Run
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/app .
CMD ["./app"]
`

	// Default project directory
	projectPath := "./aiseHi" // Modify this path as needed

	// Ensure the directory exists
	err := os.MkdirAll(projectPath, os.ModePerm)
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
}
