package configs

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func MainMenu() (string, error) {
	menu := promptui.Select{
		Label: "Select an Action",
		Items: []string{
			"Setup CI/CD",
			"Generate Dockerfile",
			"Generate Infrastructure",
			"Estimate Costs",
			"Setup Monitoring",
			"Cloud providers",
			"Fetch from github repo",
			"Exit",
		},
	}

	_, result, err := menu.Run()
	if err != nil {
		return "", fmt.Errorf("failed to select menu option: %w", err)
	}

	return result, nil
}

func InputPrompt(label string) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("failed to get input: %w", err)
	}

	return result, nil
}

func ConfirmPrompt(label string) (bool, error) {
	prompt := promptui.Prompt{
		Label:     label + " (y/n)",
		IsConfirm: true,
	}

	_, err := prompt.Run()
	if err != nil {
		return false, nil
	}

	return true, nil
}

func AWSMenu() (string, error) {
	menu := promptui.Select{
		Label: "Select an AWS Service",
		Items: []string{
			"ECR",
			"S3",
			"EKS",
			"Back to Main Menu",
		},
	}

	_, result, err := menu.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}
