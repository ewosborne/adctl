/*
Copyright © 2024 Eric Osborne
No header.
*/
package cmd

import (
	"fmt"

	"github.com/ewosborne/adctl/common"
	"github.com/spf13/cobra"
)

// disableCmd represents the disable command
var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable ad blocker. Optional duration in time.Duration format.",
	Args:  cobra.RangeArgs(0, 1),
	RunE:  DisableCmdE,
}

func DisableCmdE(cmd *cobra.Command, args []string) error {
	return printDisable(args)
}

func printDisable(args []string) error {

	var err error

	status, err := disableCommand(args)
	if err != nil {
		return err
	}

	PrintStatus(status)

	return err

}

func disableCommand(args []string) (Status, error) {

	var err error
	if len(args) == 0 {
		// handle no time duration
		err = common.AbleCommand(false, "")
	} else if len(args) == 1 {
		// handle time duration
		err = common.AbleCommand(false, args[0])
	} else {
		return Status{}, fmt.Errorf("too many arguments to disableCommand: %v", len(args))
	}

	if err != nil {
		return Status{}, err
	}

	s, err := GetStatus()

	return s, err
}

func init() {
	rootCmd.AddCommand(disableCmd)

}
