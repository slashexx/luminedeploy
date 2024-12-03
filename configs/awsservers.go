package configs

import (
	"github.com/manifoldco/promptui"
)

func AWSServerMenu() (string, error) {
	menu := promptui.Select{
		Label: "Select an AWS Region",
		Items: []string{
			"us-east-1", // US East (N. Virginia)
			"us-east-2", // US East (Ohio)
			"us-west-1", // US West (N. California)
			"us-west-2", // US West (Oregon)
			"ca-central-1", // Canada (Central)
			"sa-east-1", // South America (São Paulo)
		},
	}

	_, result, err := menu.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}
