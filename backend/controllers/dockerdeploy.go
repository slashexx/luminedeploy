package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func HandleDockerDeploy(w http.ResponseWriter, r *http.Request) {
	// Parse the request for Docker credentials and image tag
	type DockerCredentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
		ImageTag string `json:"image_tag"` // e.g., "username/repository:tag"
	}

	var creds DockerCredentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate inputs
	if creds.Username == "" || creds.Password == "" || creds.ImageTag == "" {
		http.Error(w, "Missing required fields (username, password, image_tag)", http.StatusBadRequest)
		return
	}

	// Path to the output folder
	outputPath := filepath.Join("file", "output")
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		http.Error(w, "Output folder does not exist", http.StatusInternalServerError)
		return
	}

	// Step 1: Docker login
	cmdLogin := exec.Command("docker", "login", "-u", creds.Username, "-p", creds.Password)
	cmdLogin.Dir = outputPath
	loginOutput, err := cmdLogin.CombinedOutput()
	if err != nil {
		http.Error(w, "Failed to log in to Docker Hub: "+string(loginOutput), http.StatusInternalServerError)
		return
	}

	// Step 2: Docker build
	cmdBuild := exec.Command("docker", "build", "-t", creds.ImageTag, ".")
	cmdBuild.Dir = outputPath
	buildOutput, err := cmdBuild.CombinedOutput()
	if err != nil {
		http.Error(w, "Failed to build Docker image: "+string(buildOutput), http.StatusInternalServerError)
		return
	}

	// Step 3: Docker push
	cmdPush := exec.Command("docker", "push", creds.ImageTag)
	cmdPush.Dir = outputPath
	pushOutput, err := cmdPush.CombinedOutput()
	if err != nil {
		http.Error(w, "Failed to push Docker image: "+string(pushOutput), http.StatusInternalServerError)
		return
	}

	// Step 4: Cleanup (optional)
	// Uncomment if you want to remove local images after pushing
	// cmdCleanup := exec.Command("docker", "rmi", creds.ImageTag)
	// cleanupOutput, err := cmdCleanup.CombinedOutput()
	// if err != nil {
	// 	http.Error(w, "Failed to clean up local Docker image: "+string(cleanupOutput), http.StatusInternalServerError)
	// 	return
	// }

	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Successfully built and pushed image '%s' to Docker Hub", creds.ImageTag)))
}
