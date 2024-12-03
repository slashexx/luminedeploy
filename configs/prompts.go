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
			"Generate Infrastructure",
			"Estimate Costs",
			"Setup Monitoring",
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
