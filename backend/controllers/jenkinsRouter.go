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

// manipulateJenkinsfile function can modify the Jenkinsfile if needed (placeholder for future logic)
func manipulateJenkinsfile(content string) string {
	// Example: Replace placeholders or add additional stages dynamically
	return content // Keeping it simple for now
}

// GenerateJenkinsFile is a main function for routing without parameters or return type
func GenerateJenkinsFile(w http.ResponseWriter, r *http.Request) {
	// Fixed default values for the Jenkinsfile
	params := JenkinsParams{
		PipelineName: "DefaultPipeline",
		BranchName:   "main",
		BuildCommand: "go build",
		TestCommand:  "go test ./...",
		AgentLabel:   "linux",
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
