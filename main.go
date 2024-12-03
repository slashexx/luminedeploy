package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"lumine/configs"

	"lumine/integrations/providers"
	"lumine/integrations/monitoring"
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

				
				name, _ := prompt.Run()
				prompt2 := promptui.Prompt{
					Label: "Enter address :",
				}
				dirname, _ := prompt2.Run()

				err := providers.GenerateECRConfig(name, dirname)
				if err != nil {
					fmt.Println("Error generating ECR config:", err)
				} else {
					fmt.Println("Successfully generated ECR Terraform config at ./outputs/aws/ecr")
				}

				// case "S3":
				// 	bucketName := promptui.Prompt{
				// 		Label: "Enter the name of your S3 bucket",
				// 	}
				// 	bucketName.Run()

				// 	prompt2 := promptui.Prompt{
				// 		Label: "Enter address :",
				// 	}
	
				// 	dirname, _ := prompt2.Run()
	

				// 	err := providers.GenerateS3Config(bucketName, dirname)
				// 	if err != nil {
				// 		fmt.Println("Error generating S3 config:", err)
				// 	} else {
				// 		fmt.Println("Successfully generated S3 Terraform config at ./outputs/aws/s3")
				// 	}

				case "EKS":
					prompt := promptui.Prompt{
						Label: "Enter the name of your EKS cluster",
					}
					clusterName, _ := prompt.Run()

					serverName, _ := configs.AWSServerMenu()

					prompt2 := promptui.Prompt{
						Label: "Enter the directory you want the terraform configs",
					}

					dirname, _ := prompt2.Run()

					err := providers.GenerateEKSConfig(clusterName, serverName, dirname)
					if err != nil {
						fmt.Println("Error generating EKS config:", err)
					} else {
						fmt.Println("Successfully generated EKS Terraform config at ./outputs/aws/eks")
					}
			}

		case "Setup Monitoring":
			// Prompts for setting up monitoring (e.g., Prometheus)
			monitoringChoice, err := configs.InputPrompt("Would you like to set up Prometheus monitoring? (y/n)")
			if err != nil {
				fmt.Println(configs.FormatError(err))
				break
			}

			if monitoringChoice == "y" || monitoringChoice == "Y" {
				// Set up Prometheus by calling the monitoring setup function
				err := monitoring.SetupPrometheusMonitoring()
				if err != nil {
					fmt.Println("Error setting up Prometheus:", err)
				} else {
					fmt.Println("Prometheus setup complete!")
				}
			} else {
				fmt.Println("Skipping monitoring setup.")
			}
		case "Exit":
			return
		default:
			fmt.Println("Invalid choice, returning to main menu...")
		}
	}
}