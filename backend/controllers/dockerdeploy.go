package controllers

import (
	"encoding/json"
	"fmt"
	// "github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func HandleDockerDeploy(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form (for credentials + file upload)
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Missing Docker Hub credentials", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		http.Error(w, "No files provided", http.StatusBadRequest)
		return
	}

	// Save uploaded files
	uploadDir := "./uploads"
	os.MkdirAll(uploadDir, os.ModePerm)

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Failed to open file: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		destPath := filepath.Join(uploadDir, fileHeader.Filename)
		out, err := os.Create(destPath)
		if err != nil {
			http.Error(w, "Failed to save file: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer out.Close()

		io.Copy(out, file)
	}

	// Ensure Dockerfile exists
	dockerfilePath := filepath.Join(uploadDir, "Dockerfile")
	if _, err := os.Stat(dockerfilePath); os.IsNotExist(err) {
		http.Error(w, "Dockerfile not found in uploaded files", http.StatusBadRequest)
		return
	}

	// Docker login
	loginCmd := exec.Command("docker", "login", "-u", username, "--password-stdin")
	loginStdin, err := loginCmd.StdinPipe()
	if err != nil {
		http.Error(w, "Failed to create stdin pipe for login: "+err.Error(), http.StatusInternalServerError)
		return
	}
	go func() {
		defer loginStdin.Close()
		io.WriteString(loginStdin, password)
	}()

	loginCmd.Stdout = os.Stdout
	loginCmd.Stderr = os.Stderr
	if err = loginCmd.Run(); err != nil {
		http.Error(w, "Docker login failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Build Docker image
	imageName := fmt.Sprintf("%s/custom-app:latest", username)
	buildCmd := exec.Command("docker", "build", "-t", imageName, uploadDir)
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	if err = buildCmd.Run(); err != nil {
		http.Error(w, "Docker build failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Push Docker image
	pushCmd := exec.Command("docker", "push", imageName)
	pushCmd.Stdout = os.Stdout
	pushCmd.Stderr = os.Stderr
	if err = pushCmd.Run(); err != nil {
		http.Error(w, "Docker push failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Deployment successful",
		"image":   imageName,
	})
}