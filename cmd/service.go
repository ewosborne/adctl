/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// servicesCmd represents the services command
var servicesCmd = &cobra.Command{
	Use:   "service",
	Short: "Alter filtered services",
}

func init() {
	rootCmd.AddCommand(servicesCmd)
}
