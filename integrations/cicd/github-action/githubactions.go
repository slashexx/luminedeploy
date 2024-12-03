package githubaction

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/manifoldco/promptui"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
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
          key: go-<runner_os>-go-{{ .CacheKey }}
          restore-keys: |
            go-<runner_os>-

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

// manipulateYAMLToAddRunnerOS function dynamically adds the runner.os variable to YAML
func manipulateYAMLToAddRunnerOS(yamlContent string) string {
	// Replace <runner_os> placeholder with the actual value for runner.os
	updatedYAML := strings.ReplaceAll(yamlContent, "<runner_os>", "${{ runner.os }}")

	// Return the modified YAML content
	return updatedYAML
}

// Function to prompt user input with validation
func promptInput(label string, validateFunc promptui.ValidateFunc) string {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validateFunc,
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v", err)
	}
	return result
}

func GenerateGithubActions() {
	// Prompt user for input using PromptUI
	workflowName := promptInput("Enter the Workflow Name", nil)
	triggerEvents := promptInput("Enter the Trigger Events (e.g., push, pull_request)", nil)
	goVersion := promptInput("Enter the Go Version (e.g., 1.20)", nil)
	buildCommand := promptInput("Enter the Build Command (e.g., go build)", nil)
	testCommand := promptInput("Enter the Test Command (e.g., go test ./...)", nil)

	// Example pattern for hash calculation
	hash, err := hashFiles("**/go.sum")
	if err != nil {
		log.Fatal(err)
	}

	// Set the computed hash as part of the GitHub Action parameters
	params := GitHubActionParams{
		WorkflowName:  workflowName,
		TriggerEvents: triggerEvents,
		GoVersion:     goVersion,
		BuildCommand:  buildCommand,
		TestCommand:   testCommand,
		CacheKey:      hash,
	}

	// Generate initial YAML without runner.os placeholder
	initialYAML := generateGitHubActionYAML(params)

	// Manipulate the YAML to add runner.os dynamically
	finalYAML := manipulateYAMLToAddRunnerOS(initialYAML)

	// Define the output file path
	outputFile := "github_action.yml"

	// Write the final YAML to a file
	err = ioutil.WriteFile(outputFile, []byte(finalYAML), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Confirmation message
	fmt.Printf("GitHub Action YAML has been written to: %s\n", outputFile)
}
