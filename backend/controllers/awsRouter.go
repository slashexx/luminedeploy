package controllers

import (
	"fmt"
	"os"
	"text/template"
	"net/http"
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

// GenerateECRConfig generates the ECR configuration without parameters or return type
func GenerateECRConfig(w http.ResponseWriter, r *http.Request) {
	// Hardcoded name and output directory
	name := "my-ecr-repository"
	outputDir := "./ecr-config"

	// Parse the ECR template
	tmpl, err := template.New("ecr").Parse(ecrTemplate)
	if err != nil {
		fmt.Printf("failed to parse template: %v\n", err)
		return
	}

	// Data for the template
	data := map[string]string{"Name": name}

	// Ensure the directory exists
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err = os.MkdirAll(outputDir, os.ModePerm); err != nil {
			fmt.Printf("failed to create output directory: %v\n", err)
			return
		}
	}

	// Define the file path
	filePath := fmt.Sprintf("%s/main.tf", outputDir)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	// Execute the template and write to file
	if err := tmpl.Execute(file, data); err != nil {
		fmt.Printf("failed to execute template: %v\n", err)
	}
}

// GenerateEKSConfig generates the EKS configuration without parameters or return type
func GenerateEKSConfig(w http.ResponseWriter, r *http.Request) {
	// Hardcoded cluster name, region, and output directory
	clusterName := "my-eks-cluster"
	region := "us-west-2"
	outputDir := "./eks-config"

	// Parse the EKS template
	tmpl, err := template.New("eks").Parse(eksTemplate)
	if err != nil {
		fmt.Printf("failed to parse template: %v\n", err)
		return
	}

	// Data for the template
	data := map[string]string{
		"ClusterName": clusterName,
		"Region":      region,
	}

	// Ensure the directory exists
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err = os.MkdirAll(outputDir, os.ModePerm); err != nil {
			fmt.Printf("failed to create output directory: %v\n", err)
			return
		}
	}

	// Define the file path
	filePath := fmt.Sprintf("%s/main.tf", outputDir)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	// Execute the template and write to file
	if err := tmpl.Execute(file, data); err != nil {
		fmt.Printf("failed to execute template: %v\n", err)
	}
}

// GenerateS3Config generates the S3 configuration without parameters or return type
func GenerateS3Config(w http.ResponseWriter, r *http.Request) {
	// Hardcoded bucket name and output directory
	bucketName := "my-s3-bucket"
	outputDir := "./s3-config"

	// Parse the S3 template
	tmpl, err := template.New("s3").Parse(s3Template)
	if err != nil {
		fmt.Printf("failed to parse template: %v\n", err)
		return
	}

	// Data for the template
	data := map[string]string{"BucketName": bucketName}

	// Ensure the directory exists
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err = os.MkdirAll(outputDir, os.ModePerm); err != nil {
			fmt.Printf("failed to create output directory: %v\n", err)
			return
		}
	}

	// Define the file path
	filePath := fmt.Sprintf("%s/main.tf", outputDir)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	// Execute the template and write to file
	if err := tmpl.Execute(file, data); err != nil {
		fmt.Printf("failed to execute template: %v\n", err)
	}
}