package cmd

import (
	"fmt"
	"github.com/samfrm/prepush/internal/helper"
	"github.com/spf13/cobra"
)

var fileType int

const (
	YAML = iota + 1
	YML  = iota + 1
	JSON
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a new config file",
	Long:  `Initialise a new config file if not exists in your project root.`,
	Run: func(cmd *cobra.Command, args []string) {
		e, configFile := helper.GetConfigFile()
		if !e {
			askForConfigFileType()
		} else {
			helper.PrintWarning("\nThe config file is already exists!")
			helper.PrintInfo(fmt.Sprintf("Config file path: %s\n", configFile.Name()))
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func askForConfigFileType() {
	fmt.Println("Please select the configuration file type:")
	fmt.Println("1. YAML")
	fmt.Println("2. YML")
	fmt.Println("3. JSON")
	fmt.Print("Please choose a number: ")

	// read user input
	_, err := fmt.Scanln(&fileType)
	if err != nil {
		helper.PrintWarning("Input must be an integer. Please try again.")
		return
	}

	// check user input
	if fileType < YAML || fileType > JSON {
		helper.PrintError("Invalid input. Please try again.")
		return
	}

	// create config file
	createConfigFile(fileType)
}

func createConfigFile(fileType int) {
	if fileType == YAML {
		helper.MakeConfigFile("yaml")
	} else if fileType == YML {
		helper.MakeConfigFile("yml")
	} else {
		helper.MakeConfigFile("json")
	}
}
