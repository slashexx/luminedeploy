package main

import (
	"fmt"
	"lumine/configs"
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

		case "Generate Infrastructure":
			fmt.Println("Generating Infrastructure...")

		case "Estimate Costs":
			fmt.Println("Estimating Costs...")

		case "Setup Monitoring":
			fmt.Println("Setting up Monitoring...")

		case "Exit":
			fmt.Println("Exiting...")
			return
		}
	}
}
