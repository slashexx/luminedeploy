package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"text/template"
	"net/http"
)

// JenkinsParams struct to hold input parameters for the Jenkinsfile
type JenkinsParams struct {
	PipelineName string
	BranchName   string
	BuildCommand string
	TestCommand  string
	AgentLabel   string
}

// generateJenkinsfile generates a Jenkinsfile based on the given parameters
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

// GenerateJenkinsFile handles the HTTP request and generates the Jenkinsfile based on user input
func GenerateJenkinsFile(w http.ResponseWriter, r *http.Request) {
	// Get parameters from the request (using URL query params for simplicity)
	pipelineName := r.URL.Query().Get("pipeline_name")
	branchName := r.URL.Query().Get("branch_name")
	buildCommand := r.URL.Query().Get("build_command")
	testCommand := r.URL.Query().Get("test_command")
	agentLabel := r.URL.Query().Get("agent_label")

	// Use default values if parameters are not provided
	if pipelineName == "" {
		pipelineName = "DefaultPipeline"
	}
	if branchName == "" {
		branchName = "main"
	}
	if buildCommand == "" {
		buildCommand = "go build"
	}
	if testCommand == "" {
		testCommand = "go test ./..."
	}
	if agentLabel == "" {
		agentLabel = "linux"
	}

	// Set up the parameters to be passed to the Jenkinsfile template
	params := JenkinsParams{
		PipelineName: pipelineName,
		BranchName:   branchName,
		BuildCommand: buildCommand,
		TestCommand:  testCommand,
		AgentLabel:   agentLabel,
	}

	// Generate the Jenkinsfile based on the provided parameters
	jenkinsfile := generateJenkinsfile(params)

	// Write the generated Jenkinsfile to the response
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jenkinsfile))

	// Optionally, save the Jenkinsfile to a file (for example, "Jenkinsfile")
	err := ioutil.WriteFile("Jenkinsfile", []byte(jenkinsfile), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Confirmation message
	fmt.Printf("Jenkinsfile has been written to: Jenkinsfile\n")
}
