/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var statusToggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Toggle adblocker between enabled and disabled.",
	RunE:  ToggleCmdE,
}

func ToggleCmdE(cmd *cobra.Command, args []string) error {
	return printToggle()
}

func init() {
	rootCmd.AddCommand(statusToggleCmd)
}
