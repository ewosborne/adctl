/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check and change adblocking status",
	RunE:  StatusGetCmdE,
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
