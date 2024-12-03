package configs

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

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
