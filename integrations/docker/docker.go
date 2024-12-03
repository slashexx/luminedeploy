package docker

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenerateGoDockerfile(projectPath string) error {
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

	// Ensure the directory exists
	err := os.MkdirAll(projectPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory %s: %w", projectPath, err)
	}

	// Write the Dockerfile
	dockerfilePath := filepath.Join(projectPath, "Dockerfile")
	file, err := os.Create(dockerfilePath)
	if err != nil {
		return fmt.Errorf("failed to create Dockerfile: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(dockerfileContent)
	if err != nil {
		return fmt.Errorf("failed to write Dockerfile content: %w", err)
	}

	return nil
}
