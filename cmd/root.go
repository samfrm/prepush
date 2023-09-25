package cmd

import (
	"github.com/samfrm/prepush/internal/helper"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "prepush",
	Short: "Run scripts before push!",
	Long: `Welcome to PrePush!
This tool will help you store all your scripts and todos before push your changes.`,
	Run: func(cmd *cobra.Command, args []string) {
		// run commands
		helper.RunCommands()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
