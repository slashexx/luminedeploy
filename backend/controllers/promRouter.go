package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"text/template"
	"net/http"
	// "github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// Docker Compose template
const dockerComposeTemplate = `
version: '3'
services:
  prometheus:
    image: prom/prometheus
    volumes:
      - ./prometheus.yml:/etc/prometheus/prometheus.yml
    ports:
      - "9090:9090"
`

// Prometheus config template
const prometheusConfigTemplate = `
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'node_exporter'
    static_configs:
      - targets: ['localhost:9100']
`

// Service defines the structure of a docker-compose service
type Service struct {
	Image       string            `yaml:"image"`
	Ports       []string          `yaml:"ports,omitempty"`
	Environment map[string]string `yaml:"environment,omitempty"`
}

// ComposeFile represents the structure of the docker-compose file
type ComposeFile struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
}

// AppendOrCreateDockerCompose appends to or creates a docker-compose.yml file
func AppendOrCreateDockerCompose(serviceName string, newService Service) error {
	filename := "docker-compose.yml"

	var compose ComposeFile

	// Check if the file exists
	if _, err := os.Stat(filename); err == nil {
		// Read the existing file
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("failed to read existing docker-compose.yml: %w", err)
		}

		// Parse existing YAML
		if err := yaml.Unmarshal(data, &compose); err != nil {
			return fmt.Errorf("failed to parse docker-compose.yml: %w", err)
		}
	} else {
		// Initialize a new ComposeFile if it doesn't exist
		compose = ComposeFile{
			Version:  "3.8",
			Services: make(map[string]Service),
		}
	}

	// Add or update the service
	compose.Services[serviceName] = newService

	// Marshal the updated ComposeFile back to YAML
	updatedData, err := yaml.Marshal(&compose)
	if err != nil {
		return fmt.Errorf("failed to marshal updated docker-compose.yml: %w", err)
	}

	// Write back to the file
	if err := ioutil.WriteFile(filename, updatedData, 0644); err != nil {
		return fmt.Errorf("failed to write to docker-compose.yml: %w", err)
	}

	fmt.Printf("Service '%s' added/updated successfully in docker-compose.yml\n", serviceName)
	return nil
}

// GenerateDockerCompose creates the docker-compose.yml file
func GenerateDockerCompose() error {
	// Open the docker-compose.yml file (append mode if it exists)
	file, err := os.OpenFile("docker-compose.yml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Parse the template
	tmpl, err := template.New("dockerCompose").Parse(dockerComposeTemplate)
	if err != nil {
		return err
	}

	// Write the content to the file
	return tmpl.Execute(file, nil)
}

// GeneratePrometheusConfig creates the prometheus.yml file
func GeneratePrometheusConfig() error {
	// Create the file if it doesn't exist
	file, err := os.Create("prometheus.yml")
	if err != nil {
		return err
	}
	defer file.Close()

	// Parse the template for Prometheus config
	tmpl, err := template.New("prometheusConfig").Parse(prometheusConfigTemplate)
	if err != nil {
		return err
	}

	// Write the Prometheus configuration to the file
	return tmpl.Execute(file, nil)
}

// StartPrometheus starts Prometheus using Docker Compose
func StartPrometheus() error {
	cmd := exec.Command("docker-compose", "up", "-d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// StartNodeExporter starts the Node Exporter
func StartNodeExporter() error {
	cmd := exec.Command("docker", "run", "-d", "-p", "9100:9100", "prom/node-exporter")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// StatusMessage displays the URLs where Prometheus and Node Exporter can be accessed
func StatusMessage() {
	fmt.Println("Prometheus running at http://localhost:9090")
	fmt.Println("Node Exporter running at http://localhost:9100")
}

// SetupPrometheusMonitoring handles the full setup process
func SetupPrometheusMonitoring() error {
	// Generate Docker Compose file
	err := GenerateDockerCompose()
	if err != nil {
		return fmt.Errorf("failed to generate docker-compose.yml: %v", err)
	}

	// Generate Prometheus config file
	err = GeneratePrometheusConfig()
	if err != nil {
		return fmt.Errorf("failed to generate prometheus.yml: %v", err)
	}

	// Start Prometheus and Node Exporter services
	err = StartPrometheus()
	if err != nil {
		return fmt.Errorf("failed to start Prometheus: %v", err)
	}

	err = StartNodeExporter()
	if err != nil {
		return fmt.Errorf("failed to start Node Exporter: %v", err)
	}

	// Print the status message
	StatusMessage()

	return nil
}

// PrometheusSetupHandler HTTP handler for Prometheus setup
func PrometheusSetupHandler(w http.ResponseWriter, r *http.Request) {
	err := SetupPrometheusMonitoring()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error setting up Prometheus: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message":   "Prometheus setup completed successfully",
		"prometheus": "http://localhost:9090",
		"node_exporter": "http://localhost:9100",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// InitRouter initializes the router and sets up the routes
// func InitRouter() *mux.Router {
// 	r := mux.NewRouter()

// 	// Route for Prometheus setup
// 	r.HandleFunc("/api/setup-prometheus", PrometheusSetupHandler).Methods("POST")

// 	return r
// }

// StartServer starts the HTTP server
// func StartServer() {
// 	// Initialize router
// 	r := InitRouter()

// 	// Start the HTTP server
// 	port := ":8080"
// 	fmt.Printf("Server started on port %s\n", port)
// 	if err := http.ListenAndServe(port, r); err != nil {
// 		fmt.Printf("Error starting server: %v\n", err)
// 	}
// }
