package providers

import (
	"fmt"
	"os"
	"text/template"
)


const eksTemplate = `
resource "aws_eks_cluster" "{{.ClusterName}}" {
  name     = "{{.ClusterName}}"
  role_arn = "arn:aws:iam::your_account_id:role/eks-service-role"

  vpc_config {
    subnet_ids = ["subnet-0123456789abcdef0", "subnet-abcdef0123456789"]
  }
}
`

const ecrTemplate = `
resource "aws_ecr_repository" "{{.Name}}" {
  name = "{{.Name}}"
  image_tag_mutability = "MUTABLE"
  image_scanning_configuration {
    scan_on_push = true
  }
}
`

const s3Template = `
resource "aws_s3_bucket" "{{.BucketName}}" {
  bucket = "{{.BucketName}}"
  acl    = "private"

  versioning {
    enabled = true
  }

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm = "AES256"
      }
    }
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


func GenerateEKSConfig(clusterName, region, outputDir string) error {
	tmpl, err := template.New("eks").Parse(eksTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	data := map[string]string{
		"ClusterName": clusterName,
		"Region":      region,
	}

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

func GenerateS3Config(bucketName, outputDir string) error {
	tmpl, err := template.New("s3").Parse(s3Template)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	data := map[string]string{"BucketName": bucketName}

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
