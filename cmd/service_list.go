/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// serviceListCmd represents the list command
var serviceListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
}

func init() {
	servicesCmd.AddCommand(serviceListCmd)
}
