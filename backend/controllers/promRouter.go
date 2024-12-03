package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"text/template"
	"net/http"
	// "gopkg.in/yaml.v3"
	// "io/ioutil"
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

// Default Prometheus config template
const defaultPrometheusConfigTemplate = `
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'node_exporter'
    static_configs:
      - targets: ['localhost:9100']
`

// Dynamic Prometheus config template
const dynamicPrometheusConfigTemplate = `
global:
  scrape_interval: {{.ScrapeInterval}}

scrape_configs:
  - job_name: 'node_exporter'
    static_configs:
      - targets: {{.Targets}}

  - job_name: 'custom_job'
    static_configs:
      - targets: {{.AdditionalTargets}}
`

// PrometheusConfig holds dynamic configuration data
type PrometheusConfig struct {
	ScrapeInterval    string
	Targets           []string
	AdditionalTargets []string
}

// GeneratePrometheusConfig generates a Prometheus config file dynamically or uses defaults
func GeneratePrometheusConfig(config PrometheusConfig) error {
	file, err := os.Create("prometheus.yml")
	if err != nil {
		return fmt.Errorf("failed to create prometheus.yml: %v", err)
	}
	defer file.Close()

	tmpl, err := template.New("prometheusConfig").Parse(dynamicPrometheusConfigTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse Prometheus config template: %v", err)
	}

	// Populate template with provided or default values
	if config.ScrapeInterval == "" {
		config.ScrapeInterval = "15s"
	}
	if len(config.Targets) == 0 {
		config.Targets = []string{"localhost:9100"}
	}
	if len(config.AdditionalTargets) == 0 {
		config.AdditionalTargets = []string{}
	}

	// Execute template and write to file
	return tmpl.Execute(file, config)
}

// GenerateDockerCompose creates or updates the docker-compose.yml file
func GenerateDockerCompose() error {
	file, err := os.Create("docker-compose.yml")
	if err != nil {
		return fmt.Errorf("failed to create docker-compose.yml: %v", err)
	}
	defer file.Close()

	tmpl, err := template.New("dockerCompose").Parse(dockerComposeTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse Docker Compose template: %v", err)
	}

	// Execute template and write to file
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

// StatusMessage displays Prometheus and Node Exporter URLs
func StatusMessage() {
	fmt.Println("Prometheus running at http://localhost:9090")
	fmt.Println("Node Exporter running at http://localhost:9100")
}

// SetupPrometheusMonitoring handles the complete setup process dynamically
func SetupPrometheusMonitoring(config PrometheusConfig) error {
	// Generate Docker Compose file
	err := GenerateDockerCompose()
	if err != nil {
		return fmt.Errorf("failed to generate docker-compose.yml: %v", err)
	}

	// Generate Prometheus config file dynamically
	err = GeneratePrometheusConfig(config)
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
	var requestBody map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Extract dynamic configuration values from request body
	config := PrometheusConfig{
		ScrapeInterval:    getStringValue(requestBody, "scrapeInterval", "15s"),
		Targets:           getStringArray(requestBody, "targets", []string{"localhost:9100"}),
		AdditionalTargets: getStringArray(requestBody, "additionalTargets", []string{}),
	}

	// Setup Prometheus with the dynamic configuration
	err := SetupPrometheusMonitoring(config)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error setting up Prometheus: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message":        "Prometheus setup completed successfully",
		"prometheus_url": "http://localhost:9090",
		"node_exporter_url": "http://localhost:9100",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Helper to extract a string value or return a default
func getStringValue(data map[string]interface{}, key, defaultValue string) string {
	if value, ok := data[key]; ok {
		return fmt.Sprintf("%v", value)
	}
	return defaultValue
}

// Helper to extract an array of strings or return a default
func getStringArray(data map[string]interface{}, key string, defaultValue []string) []string {
	if value, ok := data[key]; ok {
		if array, ok := value.([]interface{}); ok {
			strArray := make([]string, len(array))
			for i, v := range array {
				strArray[i] = fmt.Sprintf("%v", v)
			}
			return strArray
		}
	}
	return defaultValue
}
