package jenkins

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"text/template"

	"github.com/manifoldco/promptui"
)

// JenkinsParams struct to hold input parameters for the Jenkinsfile
type JenkinsParams struct {
	PipelineName string
	BranchName   string
	BuildCommand string
	TestCommand  string
	AgentLabel   string
}

// generateJenkinsfile generates a Jenkinsfile based on the user's input
func generateJenkinsfile(params JenkinsParams) string {
	// Define a Jenkinsfile template
	const jenkinsfileTemplate = `
pipeline {
    agent { label '{{ .AgentLabel }}' }
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage('Build') {
            steps {
                sh '{{ .BuildCommand }}'
            }
        }
        stage('Test') {
            steps {
                sh '{{ .TestCommand }}'
            }
        }
    }
    post {
        always {
            echo "Pipeline '{{ .PipelineName }}' completed for branch '{{ .BranchName }}'"
        }
    }
}
`

	// Create a buffer to hold the output Jenkinsfile
	var result bytes.Buffer

	// Parse and execute the template with provided parameters
	tmpl, err := template.New("jenkinsfile").Parse(jenkinsfileTemplate)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(&result, params)
	if err != nil {
		log.Fatal(err)
	}

	// Return the generated Jenkinsfile as a string
	return result.String()
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

// manipulateJenkinsfile function can modify the Jenkinsfile if needed (placeholder for future logic)
func manipulateJenkinsfile(content string) string {
	// Example: Replace placeholders or add additional stages dynamically
	updatedContent := strings.ReplaceAll(content, "placeholder", "dynamic_value")
	return updatedContent
}

func GenerateJenkinsFile() {
	// Prompt user for input using PromptUI
	pipelineName := promptInput("Enter the Pipeline Name", nil)
	branchName := promptInput("Enter the Branch Name (e.g., main)", nil)
	buildCommand := promptInput("Enter the Build Command (e.g., go build)", nil)
	testCommand := promptInput("Enter the Test Command (e.g., go test ./...)", nil)
	agentLabel := promptInput("Enter the Agent Label (e.g., linux, docker, etc.)", nil)

	// Set the parameters for the Jenkinsfile
	params := JenkinsParams{
		PipelineName: pipelineName,
		BranchName:   branchName,
		BuildCommand: buildCommand,
		TestCommand:  testCommand,
		AgentLabel:   agentLabel,
	}

	// Generate the Jenkinsfile
	initialJenkinsfile := generateJenkinsfile(params)

	// Manipulate the Jenkinsfile if needed
	finalJenkinsfile := manipulateJenkinsfile(initialJenkinsfile)

	// Define the output file path
	outputFile := "Jenkinsfile"

	// Write the final Jenkinsfile to a file
	err := ioutil.WriteFile(outputFile, []byte(finalJenkinsfile), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Confirmation message
	fmt.Printf("Jenkinsfile has been written to: %s\n", outputFile)
}
