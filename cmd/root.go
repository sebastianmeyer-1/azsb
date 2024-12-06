/*
Copyright © 2024 Sebastian Meyer sebastian.meyer1@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "azsb",
	Short: "Azure Service Bus CLI",
	Long: `Azure Service Bus CLI is a Toolset helping customers 
	integrating legacy systems with Azure Cloud`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
