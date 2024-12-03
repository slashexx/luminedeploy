package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow all origins (change to specific origin if needed)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func detectGoProject(files []string) (bool, string, error) {
	var hasGoMod, hasMainGo bool

	for _, file := range files {
		fileName := filepath.Base(file)
		if fileName == "go.mod" {
			hasGoMod = true
		}
		if fileName == "main.go" {
			hasMainGo = true
		}
		if hasGoMod && hasMainGo {
			return true, file, nil
		}
	}

	if !hasGoMod {
		return false, "", fmt.Errorf("missing go.mod file")
	}

	if !hasMainGo {
		return false, "", fmt.Errorf("missing main.go file")
	}

	return false, "", fmt.Errorf("incomplete Go project")
}


func main() {
	r := mux.NewRouter()

	// Route to handle file uploads
	r.HandleFunc("/upload", uploadHandler).Methods("POST")

	// Start the server
	fmt.Println("Server running on :8080")
	handler := enableCORS(r)
	http.ListenAndServe(":8080", handler)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    // Ensure uploads directory exists
    uploadDir := "./uploads"
    if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
        fmt.Println("Error creating uploads directory:", err)
        http.Error(w, "Failed to create uploads directory", http.StatusInternalServerError)
        return
    }

    // Parse multipart form
    if err := r.ParseMultipartForm(10 << 20); err != nil {
        fmt.Println("Error parsing form:", err)
        http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
        return
    }

    files := r.MultipartForm.File["file"]
    if len(files) == 0 {
        fmt.Println("No files found in request")
        http.Error(w, "No files uploaded", http.StatusBadRequest)
        return
    }

    for _, fileHeader := range files {
        file, err := fileHeader.Open()
        if err != nil {
            fmt.Println("Error opening file:", err)
            http.Error(w, "Failed to read file: "+err.Error(), http.StatusBadRequest)
            return
        }
        defer file.Close()

        // Save uploaded file
        filePath := filepath.Join(uploadDir, fileHeader.Filename)
        outFile, err := os.Create(filePath)
        if err != nil {
            fmt.Println("Error creating file:", err)
            http.Error(w, "Failed to save file: "+err.Error(), http.StatusInternalServerError)
            return
        }
        defer outFile.Close()

        if _, err = io.Copy(outFile, file); err != nil {
            fmt.Println("Error writing file:", err)
            http.Error(w, "Failed to save file: "+err.Error(), http.StatusInternalServerError)
            return
        }

        // Optionally unzip if file is a ZIP archive
        if filepath.Ext(filePath) == ".zip" {
            extractDir := filepath.Join(uploadDir, fileHeader.Filename+"_extracted")
            if err := unzip(filePath, extractDir); err != nil {
                fmt.Println("Error unzipping file:", err)
                http.Error(w, "Failed to extract file: "+err.Error(), http.StatusInternalServerError)
                return
            }
        }
    }

    fmt.Println("All files uploaded and processed successfully")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Files uploaded successfully"))
}

func unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, file := range r.File {
		filePath := filepath.Join(dest, file.Name)

		// Prevent Zip Slip vulnerability
		if !filepath.HasPrefix(filePath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("invalid file path: %s", filePath)
		}

		if file.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer outFile.Close()

		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		_, err = io.Copy(outFile, rc)
		if err != nil {
			return err
		}
	}
	return nil
}

func zipFiles(files []string, zipPath string) error {
    newZipFile, err := os.Create(zipPath)
    if err != nil {
        return err
    }
    defer newZipFile.Close()

    zipWriter := zip.NewWriter(newZipFile)
    defer zipWriter.Close()

    for _, file := range files {
        fileToZip, err := os.Open(file)
        if err != nil {
            return err
        }
        defer fileToZip.Close()

        writer, err := zipWriter.Create(filepath.Base(file))
        if err != nil {
            return err
        }

        _, err = io.Copy(writer, fileToZip)
        if err != nil {
            return err
        }
    }

    return nil
}



func generateGoDockerfile(projectPath string) (string, error) {
    dockerfileContent := `
# Use the official Golang image to build the Go binary
FROM golang:1.20-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod tidy

# Copy the entire project into the container
COPY . .

# Build the Go binary
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest  

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary file from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the Go binary
CMD ["./main"]
    `
    dockerfilePath := filepath.Join(projectPath, "Dockerfile")
    err := os.WriteFile(dockerfilePath, []byte(dockerfileContent), 0644)
    if err != nil {
        return "", err
    }
    return dockerfilePath, nil
}


func generateKubernetesManifests(projectPath string, imageName string) (string, string, error) {
    // Deployment manifest content
    deployment := fmt.Sprintf(`
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: go-app
  template:
    metadata:
      labels:
        app: go-app
    spec:
      containers:
        - name: go-app
          image: %s
          ports:
            - containerPort: 8080
    `, imageName)

    // Service manifest content
    service := `
apiVersion: v1
kind: Service
metadata:
  name: go-app-service
spec:
  selector:
    app: go-app
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
  type: LoadBalancer
    `
    deploymentPath := filepath.Join(projectPath, "deployment.yaml")
    servicePath := filepath.Join(projectPath, "service.yaml")

    err := os.WriteFile(deploymentPath, []byte(deployment), 0644)
    if err != nil {
        return "", "", err
    }

    err = os.WriteFile(servicePath, []byte(service), 0644)
    if err != nil {
        return "", "", err
    }

    return deploymentPath, servicePath, nil
}
