package controllers

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	// "strings"
	"text/template"
)

// GitHubActionParams struct to hold input parameters for the GitHub Action
type GitHubActionParams struct {
	WorkflowName  string
	TriggerEvents string
	GoVersion     string
	BuildCommand  string
	TestCommand   string
	CacheKey      string
}

// hashFiles computes the SHA256 hash of the contents of files matching the provided glob pattern.
func hashFiles(pattern string) (string, error) {
	var fileContents []byte

	// Walk through the files that match the pattern
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Match the file pattern (in this case, files like go.sum)
		if match, _ := filepath.Match(pattern, path); match && !info.IsDir() {
			// Read the file contents and append to the fileContents slice
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			fileContents = append(fileContents, content...)
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	// Create a SHA256 hash of the combined file contents
	hash := sha256.New()
	hash.Write(fileContents)

	// Return the hexadecimal representation of the hash
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// generateGitHubActionYAML generates a GitHub Actions YAML without runner.os
func generateGitHubActionYAML(params GitHubActionParams) string {
	// Define a template without using runner.os
	const yamlTemplate = ` 
name: {{ .WorkflowName }}

on:
  {{ .TriggerEvents }}:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest  # Static value for now

    steps:
      - name: Checkout code
        uses: actions/checkout@v2

      - name: Set up Go {{ .GoVersion }}
        uses: actions/setup-go@v3
        with:
          go-version: {{ .GoVersion }}

      - name: Cache Go Modules
        uses: actions/cache@v2
        with:
          path: ~/go/pkg/mod
          key: go-${{ runner.os }}-go-{{ .CacheKey }}
          restore-keys: |
            go-${{ runner.os }}-

      - name: Install dependencies
        run: go mod tidy

      - name: Build Project
        run: {{ .BuildCommand }}

      - name: Run Tests
        run: {{ .TestCommand }}
`
	// Create a buffer to hold the output YAML
	var result bytes.Buffer

	// Parse and execute the template with provided parameters
	tmpl, err := template.New("githubAction").Parse(yamlTemplate)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(&result, params)
	if err != nil {
		log.Fatal(err)
	}

	// Return the generated YAML as a string
	return result.String()
}

// GitHubActionHandler will be the HTTP handler function that generates the GitHub Actions YAML
func GitHubActionHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the incoming JSON request body
	var params GitHubActionParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	// Example pattern for hash calculation
	hash, err := hashFiles("**/go.sum")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error calculating hash: %v", err), http.StatusInternalServerError)
		return
	}

	// Set the computed hash as part of the GitHub Action parameters
	params.CacheKey = hash

	// Generate the GitHub Actions YAML
	finalYAML := generateGitHubActionYAML(params)

	// Write the final YAML to the response
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(finalYAML))

	// Optionally, you can save the YAML to a file (e.g., github_action.yml)
	err = ioutil.WriteFile("github_action.yml", []byte(finalYAML), 0644)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error writing file: %v", err), http.StatusInternalServerError)
		return
	}

	// Confirmation message
	fmt.Printf("GitHub Action YAML has been generated and written to github_action.yml\n")
}
