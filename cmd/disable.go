/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ewosborne/adctl/common"
	"github.com/spf13/cobra"
)

func disableCommand(dTime DisableTime) (Status, error) {

	var err error

	if dTime.HasTimeout {
		err = common.AbleCommand(false, dTime.Duration)
	} else {
		err = common.AbleCommand(false, "")
	}

	if err != nil {
		return Status{}, err
	}

	s, err := GetStatus()

	return s, err
}

var statusDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable ad blocker. Optional duration in time.Duration format.",
	Args:  cobra.RangeArgs(0, 1),
	RunE:  StatusDisableCmdE,
}

func init() {
	rootCmd.AddCommand(statusDisableCmd)
}

func StatusDisableCmdE(cmd *cobra.Command, args []string) error {

	var dTime = DisableTime{}

	switch len(args) {
	case 0:
		dTime.HasTimeout = false
	case 1:
		dTime.HasTimeout = true
		dTime.Duration = args[0]
	default:
		return fmt.Errorf("only one arg allowed for disable")
	}

	return printDisable(dTime)
}

func printDisable(dTime DisableTime) error {

	var err error

	status, err := disableCommand(dTime)
	if err != nil {
		return err
	}

	PrintStatus(status)

	return err

}
