package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"lumine/configs"

	"lumine/integrations/providers"
)

func main() {

	const (
		Logo = `
██╗     ██╗   ██╗███╗   ███╗██╗███╗   ██╗███████╗
██║     ██║   ██║████╗ ████║██║████╗  ██║██╔════╝
██║     ██║   ██║██╔████╔██║██║██╔██╗ ██║█████╗  
██║     ██║   ██║██║╚██╔╝██║██║██║╚██╗██║██╔══╝  
███████╗╚██████╔╝██║ ╚═╝ ██║██║██║ ╚████║███████╗
╚══════╝ ╚═════╝ ╚═╝     ╚═╝╚═╝╚═╝  ╚═══╝╚══════╝
                                                
	`
	)
	fmt.Println(Logo)

	for {
		choice, err := configs.MainMenu()
		if err != nil {
			fmt.Println(configs.FormatError(err))
			break
		}

		switch choice {
		case "Setup CI/CD":
			fmt.Println("Setting up CI/CD...")

		case "Cloud providers":
			cloudChoice, err := configs.AWSMenu()
			if err != nil {
				fmt.Println(configs.FormatError(err))
				break
			}

			switch cloudChoice {
			case "ECR":
				prompt := promptui.Prompt{
					Label: "Enter the name of your ECR repository",
				}

				prompt2 := promptui.Prompt{
					Label: "Enter address :",
				}

				name, _ := prompt.Run()
				dirname, _ := prompt2.Run()

				err := providers.GenerateECRConfig(name, dirname)
				if err != nil {
					fmt.Println("Error generating ECR config:", err)
				} else {
					fmt.Println("Successfully generated ECR Terraform config at ./outputs/aws/ecr")
				}

				// case "S3":
				// 	bucketName, _ := promptui.Prompt{
				// 		Label: "Enter the name of your S3 bucket",
				// 	}.Run()

				// 	err := providers.GenerateS3Config(bucketName, "./outputs/aws/s3")
				// 	if err != nil {
				// 		fmt.Println("Error generating S3 config:", err)
				// 	} else {
				// 		fmt.Println("Successfully generated S3 Terraform config at ./outputs/aws/s3")
				// 	}

				// case "EKS":
				// 	clusterName, _ := promptui.Prompt{
				// 		Label: "Enter the name of your EKS cluster",
				// 	}.Run()

				// 	err := providers.GenerateEKSConfig(clusterName, "us-east-1", "./outputs/aws/eks")
				// 	if err != nil {
				// 		fmt.Println("Error generating EKS config:", err)
				// 	} else {
				// 		fmt.Println("Successfully generated EKS Terraform config at ./outputs/aws/eks")
				// 	}
			}
		default:
			fmt.Println("Invalid choice, returning to main menu...")
		}
	}
}
