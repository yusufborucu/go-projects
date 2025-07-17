package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gopassgen",
	Short: "Gopassgen is a CLI application for generating secure passwords",
	Long:  "Gopassgen is a command-line application that allows you to generate secure, random passwords based on your specific needs. It is written with the Cobra library.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gopassgen CLI - use --help for the command")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
