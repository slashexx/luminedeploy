package providers

import (
	"fmt"
	"os"
	"text/template"
)

const ecrTemplate = `
resource "aws_ecr_repository" "{{.Name}}" {
  name = "{{.Name}}"
  image_tag_mutability = "MUTABLE"
  image_scanning_configuration {
    scan_on_push = true
  }
}
`

func GenerateECRConfig(name, outputDir string) error {
	tmpl, err := template.New("ecr").Parse(ecrTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	data := map[string]string{"Name": name}

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err = os.MkdirAll(outputDir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	}

	filePath := fmt.Sprintf("%s/main.tf", outputDir)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}
