/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ewosborne/adctl/common"
	"github.com/spf13/cobra"
)

var statusEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable ad blocking",
	RunE:  StatusEnableCmdE,
}

func StatusEnableCmdE(cmd *cobra.Command, flags []string) error {
	return printEnable()
}

func init() {
	rootCmd.AddCommand(statusEnableCmd)

}

func printEnable() error {

	var err error
	status, err := enableCommand()
	if err != nil {
		return err
	}
	err = PrintStatus(status)
	if err != nil {
		return err
	}
	return nil
}

func enableCommand() (Status, error) {

	err := common.AbleCommand(true, "")
	if err != nil {
		return Status{}, err
	}

	status, err := GetStatus()
	if err != nil {
		return Status{}, err
	}

	return status, nil
}
