package monitoring

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
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

// GenerateDockerCompose will append to docker-compose.yml if it exists, or create it if it doesn't.
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

// GeneratePrometheusConfig creates prometheus.yml if it doesn't exist.
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

// StartPrometheus starts Prometheus using Docker Compose.
func StartPrometheus() error {
	cmd := exec.Command("docker-compose", "up", "-d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// StartNodeExporter starts the Node Exporter.
func StartNodeExporter() error {
	cmd := exec.Command("docker", "run", "-d", "-p", "9100:9100", "prom/node-exporter")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// StatusMessage displays the URLs where Prometheus and Node Exporter can be accessed.
func StatusMessage() {
	fmt.Println("Prometheus running at http://localhost:9090")
	fmt.Println("Node Exporter running at http://localhost:9100")
}
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
