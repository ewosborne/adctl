/*
Copyright © 2024 Eric Osborne
No header.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// toggleCmd represents the toggle command
var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Toggle adblocker between enabled and disabled.",
	RunE:  ToggleCmdE,
}

func ToggleCmdE(cmd *cobra.Command, args []string) error {
	return printToggle()
}

func printToggle() error {
	var err error

	err = toggleCommand()
	if err != nil {
		return err
	}

	status, err := GetStatus()
	if err != nil {
		return err
	}
	PrintStatus(status)
	return nil
}

func toggleCommand() error {
	status, err := GetStatus()
	if err != nil {
		return err
	}

	dTime := DisableTime{HasTimeout: false}
	switch status.Protection_enabled {
	case true:
		disableCommand(dTime)
	case false:
		enableCommand()
	}

	return nil
}

func init() {
	rootCmd.AddCommand(toggleCmd)
}
