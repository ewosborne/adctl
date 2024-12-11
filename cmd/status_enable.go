/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ewosborne/adctl/common"
	"github.com/spf13/cobra"
)

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable ad blocking",
	RunE:  EnableCmdE,
}

func EnableCmdE(cmd *cobra.Command, flags []string) error {
	return printEnable()
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

func init() {
	statusCmd.AddCommand(enableCmd)

}
